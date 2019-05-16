package main

/**
 *  Composite Key functions
 **/

import (
        "encoding/json"
        "fmt"
        "github.com/hyperledger/fabric/core/chaincode/shim"
        "github.com/hyperledger/fabric/protos/ledger/queryresult"
        "github.com/hyperledger/fabric/protos/peer"
        "strconv"
        "strings"
)

// Object type constant
const objectType = "ParentCompany~CurrentCompany~EmployeeType"

func (contract *ContractChaincode) AddEmployeeWithCompositeKey(stub shim.ChaincodeStubInterface) peer.Response {
        value := []byte{0x00}
        stub.PutState(objectType, value)
        addData(stub, 10, "ov", 10000, 9.30, "permanent", "hitachi", "hitachi")
        addData(stub, 11, "karthik", 10000, 9.30, "permanent", "hitachi", "hitachi")
        addData(stub, 12, "nirosh", 10000, 9.30, "permanent", "hitachi", "hitachi")
        addData(stub, 13, "hary", 10000, 9.30, "contract", "human", "hitachi")
        addData(stub, 14, "pro", 10000, 9.30, "contract", "human", "hitachi")
        addData(stub, 15, "kamal", 10000, 9.30, "permanent", "human", "human")
        addData(stub, 16, "surya", 10000, 9.30, "permanent", "human", "human")
        addData(stub, 17, "gons", 10000, 9.30, "permanent", "hitachi", "hitachi")
        addData(stub, 18, "san", 10000, 9.30, "contract", "human", "hitachi")
        addData(stub, 19, "maanu", 10000, 9.30, "contract", "human", "hitachi")
        addData(stub, 20, "raj", 10000, 9.30, "contract", "human", "human")
        addData(stub, 21, "vijay", 10000, 9.30, "contract", "human", "human")
        addData(stub, 22, "dhoni", 10000, 9.30, "contract", "human", "human")

        fmt.Println("initialized with sample Data")
        payload := []byte("Employee created successfully")
        return shim.Success(payload)

}

func addData(stub shim.ChaincodeStubInterface, eId int, eName string, salary int, time float64, eType, parentCompany, currentCompany string) {

        var employee = Employee{EmployeeId: eId, EmployeeName: eName, Salary: salary, WorkingHours: time, EmployeeType: eType, ParentCompany: parentCompany, CurrentCompany: currentCompany}

        jsonEmployee, _ := json.Marshal(employee)
        balanceIndexKey, _ := stub.CreateCompositeKey(objectType, []string{employee.ParentCompany, employee.CurrentCompany, employee.EmployeeType})
        stub.PutState(balanceIndexKey, jsonEmployee)
}

func (contract *ContractChaincode) GetEmployeesByPartialCompositeKey(stub shim.ChaincodeStubInterface, args []string) peer.Response {

        fmt.Printf("==== Exec qry with:  ")
        fmt.Println(args)

        // Gets the state by partial query key
        QryIterator, err := stub.GetStateByPartialCompositeKey(objectType, args)
        if err != nil {
                fmt.Printf("Error in getting by range=" + err.Error())
                return shim.Error(err.Error())
        }
        var resultJSON = "["
        counter := 0
        // Iterate to read the keys returned
        for QryIterator.HasNext() {
                // Hold pointer to the query result
                var resultKV *queryresult.KV
                var err error

                // Get the next element
                resultKV, err = QryIterator.Next()
                if err != nil {
                        fmt.Println("Err=" + err.Error())
                        return shim.Error(err.Error())
                }

                // Split the composite key and send it as part of the result set
                key, arr, _ := stub.SplitCompositeKey(resultKV.GetKey())
                fmt.Println(key)
                resultJSON += " [" + strings.Join(arr, "~") + "] "
                counter++

        }
        // Closing
        QryIterator.Close()

        resultJSON += "]"
        resultJSON = "Counter=" + strconv.Itoa(counter) + "  " + resultJSON
        fmt.Println("Done.")
        return shim.Success([]byte(resultJSON))

}
