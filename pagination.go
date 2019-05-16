package main

/**
 *  Range functions with pagination
 **/

import (
        "fmt"
        "github.com/hyperledger/fabric/core/chaincode/shim"
        "github.com/hyperledger/fabric/protos/ledger/queryresult"
        "github.com/hyperledger/fabric/protos/peer"
        "strconv"
)

/*
 * func name   : GetEmployeesByRangeWithPagination
 * description : This function return  employee details with pagination concept
 */
func (contract *ContractChaincode) GetEmployeesByRangeWithPagination(stub shim.ChaincodeStubInterface, args []string) peer.Response {

        // Check the number of arguments
        // startKey = arg[0]  endKey = arg[1]   pagesize = arg[2]
        if len(args) < 3 {
                return shim.Error("MUST provide start, end Key & Page size!!")
        }

        pagesize, _ := strconv.ParseInt(string(args[2]), 10, 32)
        bookmark := ""
        var counter = 0
        var pageCounter = 0
        var resultJSON = "["
        var hasMorePages = true

        // variables to hold query iterator and metadata
        var qryIterator shim.StateQueryIteratorInterface
        var queryMetaData *peer.QueryResponseMetadata

        var err error

        for hasMorePages {
                // Execute stub API to get the range with pagination
                qryIterator, queryMetaData, err = stub.GetStateByRangeWithPagination(args[0], args[1], int32(pagesize), bookmark)
                if err != nil {
                        fmt.Printf("Error=" + err.Error())
                        return shim.Error(err.Error())
                }

                var arr = "["
                var resultKV *queryresult.KV
                // Check if there are any more records
                for qryIterator.HasNext() {

                        // Get the next element
                        resultKV, err = qryIterator.Next()

                        // Increment Counter
                        counter++
                        if arr != "[" {
                                arr += ","
                        }
                        value := string(resultKV.GetValue())
                        arr += "\"" + resultKV.GetKey() + "\"" + " : " + value

                }
                arr += "]"
                // Increment Page Counter
                pageCounter++

                if resultJSON != "[" {
                        resultJSON += ","
                }

                resultJSON += "{\"page\":" + strconv.Itoa(pageCounter) + ",\"keys\":" + arr + "}"

                // Get start key for the next page
                bookmark = queryMetaData.Bookmark

                // boomark = bland indicates not more records
                hasMorePages = (bookmark != "")

                fmt.Printf("Page: %d   Bookmark: %s \n", pageCounter, bookmark)

                // Close the iterator
                qryIterator.Close()
        }

        resultJSON += "]"

        resultJSON = "{\"count\":" + strconv.Itoa(counter) + ",\"pages\":" + resultJSON + "}"

        return shim.Success([]byte(resultJSON))
}
