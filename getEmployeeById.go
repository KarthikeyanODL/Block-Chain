package main

/**
 * Shows how to use the "GetQueryResult" function
 **/

import (
        "fmt"
        "github.com/hyperledger/fabric/core/chaincode/shim"
        "github.com/hyperledger/fabric/protos/ledger/queryresult"
        "github.com/hyperledger/fabric/protos/peer"
        //      "strconv"
)

func (contract *ContractChaincode) GetEmployeeById(stub shim.ChaincodeStubInterface, args []string) peer.Response {

        // Query JSON received as argument
        empId := args[0]
        var result string
        qry := `{
                "selector": {
                   "employeeId": {
                          "$eq": `
        qry += "" + empId + ""
        qry += `}
                }
         }`

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

                // Increment the counter
                counter++
                key := resultKV.GetKey()
                result = string(resultKV.GetValue())

                // Print the receieved result on the console
                fmt.Printf("Result# %d   %s   %s \n\n", counter, key, result)

        }

        // Close the iterator
        QryIterator.Close()

        // Return the value
        //total := "Count=" + strconv.Itoa(counter)
        if result == "" {
                return shim.Error("invalid employeeId")
        }
        return shim.Success([]byte(result))
}
