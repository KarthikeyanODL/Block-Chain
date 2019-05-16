package main

/**
 * Shows how to use the "GetQueryResult" function with user query
 **/

import (
        "bytes"
        "fmt"
        "github.com/hyperledger/fabric/core/chaincode/shim"
        "github.com/hyperledger/fabric/protos/ledger/queryresult"
        "github.com/hyperledger/fabric/protos/peer"
        //      "strconv"
)

func (contract *ContractChaincode) ExecuteUserQuery(stub shim.ChaincodeStubInterface, args []string) peer.Response {

        // Query JSON received as argument
        qry := args[0]
        //var result string

        // set buffer to return proper json result

        var buffer bytes.Buffer
        buffer.WriteString("[")
        bArrayMemberAlreadyWritten := false

        // Print the received query on the console
        fmt.Printf("Query JSON=%s \n\n", qry)

        // GetQueryResult
        QryIterator, err := stub.GetQueryResult(qry)

        // Return if there is an error
        if err != nil {
                fmt.Println(err.Error())
                return shim.Error("Error: " + err.Error())
        }

        // Iterate through the result set
        counter := 0
        for QryIterator.HasNext() {
                // Hold pointer to the query result
                var resultKV *queryresult.KV
                var err error

                // Get the next element
                resultKV, err = QryIterator.Next()

                // Return if there is an error
                if err != nil {
                        fmt.Println("Err=" + err.Error())
                        return shim.Error("Error: " + err.Error())
                }
                // Add a comma before array members, suppress it for the first array member
                if bArrayMemberAlreadyWritten == true {
                        buffer.WriteString(",")
                }

                buffer.WriteString("{\"Key\":")
                buffer.WriteString("\"")
                buffer.WriteString(resultKV.GetKey())
                buffer.WriteString("\"")

                buffer.WriteString(", \"Record\":")
                // Record is a JSON object, so we write as-is
                buffer.WriteString(string(resultKV.GetValue()))
                buffer.WriteString("}")
                bArrayMemberAlreadyWritten = true

                // Increment the counter
                counter++
                //key := resultKV.GetKey()
                //result = string(resultKV.GetValue())

                // Print the receieved result on the console
                //fmt.Printf("Result# %d   %s   %s \n\n", counter, key, result)

        }

        buffer.WriteString("]")
        // Close the iterator
        QryIterator.Close()

        result := buffer.String()
        // Return the value
        if result == "" {
                return shim.Success([]byte("No Data found"))
        }
        return shim.Success([]byte(result))
}
