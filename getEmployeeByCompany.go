package main

/**
 * Shows how to use the "GetQueryResultWithPagination" function
 **/

import (
        "fmt"
        "github.com/hyperledger/fabric/core/chaincode/shim"
        "github.com/hyperledger/fabric/protos/ledger/queryresult"
        "github.com/hyperledger/fabric/protos/peer"
        "strconv"
)

func (contract *ContractChaincode) GetEmployeeByCompany(stub shim.ChaincodeStubInterface, args []string) peer.Response {

        var pagesize int32 = 20
        bookmark := ""
        var counter = 0
        var pageCounter = 0
        var resultJSON = "["
        var hasMorePages = true

        // Query JSON received as argument
        //company := args[0]
        //var result string

        query := args[0]
        //query := `{
        //        "selector": {
        //           "parentCompany": {
        //                  "$eq": `
        //query += args[0]
        //query += `}
        //        }
        // }`

        // variables to hold query iterator and metadata
        var qryIterator shim.StateQueryIteratorInterface
        var queryMetaData *peer.QueryResponseMetadata
        var err error
        // start the pagination read loop
        lastBookmark := ""
        for hasMorePages {
                // execute the rich query
                qryIterator, queryMetaData, err = stub.GetQueryResultWithPagination(query, pagesize, bookmark)
                if err != nil {
                        fmt.Printf("GetQueryResultWithPagination Error=" + err.Error())
                        return shim.Error(err.Error())
                }
                var arr = "["
                var resultKV *queryresult.KV
                // Result read loop only if we received a different bookmark
                if lastBookmark != queryMetaData.Bookmark {
                        for qryIterator.HasNext() {

                                // Get the next element
                                resultKV, err = qryIterator.Next()

                                // Increment Counter
                                counter++
                                if arr != "[" {
                                        arr += ","
                                }
                                arr += "\"" + string(resultKV.GetValue()) + "\""
                        }
                        arr += "]"

                        // Increment Page Counter
                        pageCounter++

                        if resultJSON != "[" {
                                resultJSON += ","
                        }

                        fmt.Printf("Page: %d \n", pageCounter)

                        resultJSON += "{\"page\":" + strconv.Itoa(pageCounter) + ",\"Record\":" + arr + "}"
                }

                // Get start key for the next page
                bookmark = queryMetaData.Bookmark

                // boomark = blank indicates no more records
                hasMorePages = (bookmark != "" && lastBookmark != bookmark)
                lastBookmark = bookmark

                // Close the iterator
                qryIterator.Close()
        }

        resultJSON += "]"
        resultJSON = "{\"count\":" + strconv.Itoa(counter) + ",\"pages\":" + resultJSON + "}"

        return shim.Success([]byte(resultJSON))
}
