/*
 * The Employee smart contract for documentation topic:
 * Writing First Blockchain Application
 * Concepts included :
 * 1. Transaction with StateData
 * 2. Endorsement policy
 * 3. Private Data Collection
 * 4. Transaction History
 * 5. Query - page, composite key
 * 6. Rich Query
 *
 * To-Do:
 * 7. Access and authorization control
 * 8. deploy in multi network
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */

import (
        "bytes"
        "encoding/json"
        "fmt"
        "github.com/hyperledger/fabric/core/chaincode/shim"
        "github.com/hyperledger/fabric/protos/ledger/queryresult"
        "github.com/hyperledger/fabric/protos/peer"
        "strconv"
)

type ContractChaincode struct {
}

/* Define the employee structure, with 7 properties
 * Different Data types used
 */
type Employee struct {
        EmployeeId     int     `json:"employeeId"`
        EmployeeName   string  `json:"employeeName"`
        Salary         int     `json:"salary"`
        WorkingHours   float64 `json:"workingHours"`
        EmployeeType   string  `json:"employeeType"`
        ParentCompany  string  `json:"parentCompany"`
        CurrentCompany string  `json:"currentCompany"`
}

/*
 * The Init method is called when the Chain code "employee contract"
 * is instantiated by the blockchain network
 */
func (contract *ContractChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
        fmt.Println("Init executed")
        return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Chain code "employee contract"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (contract *ContractChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

        // Retrieve the requested Smart Contract function and arguments
        function, args := stub.GetFunctionAndParameters()
        fmt.Println("Invoke is running " + function)

        // Route to the appropriate handler function to interact with the ledger appropriately
        if function == "addEmployee" {
                return contract.addEmployee(stub)
        } else if function == "createEmployee" {
                return contract.createEmployee(stub, args)
        } else if function == "getEmployees" {
                return contract.getEmployees(stub)
        } else if function == "sendEmployee" {
                return contract.sendEmployee(stub, args)
        } else if function == "SetSalary" {
                return contract.SetSalary(stub, args)
        } else if function == "GetSalary" {
                return contract.GetSalary(stub, args)
        } else if function == "getEmployeeTxHistory" {
                return contract.getEmployeeTxHistory(stub, args)
        } else if function == "GetEmployeesByRangeWithPagination" {
                return contract.GetEmployeesByRangeWithPagination(stub, args)
        } else if function == "AddEmployeeWithCompositeKey" {
                return contract.AddEmployeeWithCompositeKey(stub)
        } else if function == "GetEmployeesByPartialCompositeKey" {
                return contract.GetEmployeesByPartialCompositeKey(stub, args)
        } else if function == "GetEmployeeById" {
                return contract.GetEmployeeById(stub, args)
        } else if function == "ExecuteUserQuery" {
                return contract.ExecuteUserQuery(stub, args)
        } else if function == "GetEmployeeByCompany" {
                return contract.GetEmployeeByCompany(stub, args)
        }

        return shim.Error("invalid function name")
        // return peer.Response{Status:401, Message: "unAuthorized", Payload: payload}
}

/*
 * func name   : addEmployee
 * description : This function add some hardcoded employee details to the ledger
 */
func (contract *ContractChaincode) addEmployee(stub shim.ChaincodeStubInterface) peer.Response {
        employees := []Employee{
                Employee{EmployeeId: 1, EmployeeName: "karthik", Salary: 5000, WorkingHours: 9.30, EmployeeType: "contract", ParentCompany: "human", CurrentCompany: "human"},
                Employee{EmployeeId: 2, EmployeeName: "mithun", Salary: 4000, WorkingHours: 9.30, EmployeeType: "contract", ParentCompany: "human", CurrentCompany: "human"},
                Employee{EmployeeId: 3, EmployeeName: "kawakami", Salary: 10000, WorkingHours: 9.30, EmployeeType: "permanent", ParentCompany: "human", CurrentCompany: "human"},
                Employee{EmployeeId: 4, EmployeeName: "ozawa", Salary: 15000, WorkingHours: 8.30, EmployeeType: "permanent", ParentCompany: "hitachi", CurrentCompany: "hitachi"},
                Employee{EmployeeId: 5, EmployeeName: "sakura", Salary: 10000, WorkingHours: 8.30, EmployeeType: "permanent", ParentCompany: "hitachi", CurrentCompany: "hitachi"},
        }

        i := 0
        for i < len(employees) {
                employeeAsBytes, _ := json.Marshal(employees[i])
                stub.PutState(strconv.Itoa(employees[i].EmployeeId), employeeAsBytes)
                //stub.PutState("empId"+strconv.Itoa(i), employeeAsBytes)
                fmt.Println("Added", employees[i])
                i = i + 1
        }

        payload := []byte("Employee details added successfully , count: " + strconv.Itoa(i))
        return shim.Success(payload)

}

/*
 * func name   : createEmployee
 * description : Add employee details at the run time
 */
func (contract *ContractChaincode) createEmployee(stub shim.ChaincodeStubInterface, args []string) peer.Response {

        if len(args) != 7 {
                return shim.Error("Incorrect number of arguments, required: 7")
        }

        //"createEmployee","77","ov","10000","9.30","permanent","hitachi","nissan"]

        key := args[0]
        employeeId, _ := strconv.Atoi(key)
        salary, _ := strconv.Atoi(args[2])
        workingHours, _ := strconv.ParseFloat(args[3], 64)

        var employee = Employee{EmployeeId: employeeId, EmployeeName: args[1], Salary: salary, WorkingHours: workingHours, EmployeeType: args[4], ParentCompany: args[5], CurrentCompany: args[6]}

        employeeAsBytes, _ := json.Marshal(employee)
        //stub.PutState("emp-id"+strconv.Itoa(employee.employeeId), employeeAsBytes)
        stub.PutState(key, employeeAsBytes)
        fmt.Println("Created ", employee)
        payload := []byte("Employee created successfully")
        return shim.Success(payload)
}

/*
 * func name   : getEmployees
 * description : Query all employee details using range
 */

func (contract *ContractChaincode) getEmployees(stub shim.ChaincodeStubInterface) peer.Response {

        startKey := "0"
        endKey := "999"

        resultsIterator, err := stub.GetStateByRange(startKey, endKey)

        if err != nil {
                return shim.Error(err.Error())
        }
        defer resultsIterator.Close()

        var buffer bytes.Buffer
        buffer.WriteString("[")
        bArrayMemberAlreadyWritten := false

        for resultsIterator.HasNext() {
                queryResponse, err := resultsIterator.Next()
                if err != nil {
                        return shim.Error(err.Error())
                }
                // Add a comma before array members, suppress it for the first array member
                if bArrayMemberAlreadyWritten == true {
                        buffer.WriteString(",")
                }
                buffer.WriteString("{\"Key\":")
                buffer.WriteString("\"")
                buffer.WriteString(queryResponse.Key)
                buffer.WriteString("\"")

                buffer.WriteString(", \"Record\":")
                // Record is a JSON object, so we write as-is
                buffer.WriteString(string(queryResponse.Value))
                buffer.WriteString("}")
                bArrayMemberAlreadyWritten = true

        }

        buffer.WriteString("]")

        fmt.Printf("All employees:\n%s\n", buffer.String())
        result := buffer.String()
        payload := []byte(result)
        return peer.Response{Status: 200, Message: "Employee details Queried", Payload: payload}
}

/*
 * func name   : sendEmployee
 * description : Transfer contract employee to another organisation
 *               This function also emit event-status when we call this function
 */

func (contract *ContractChaincode) sendEmployee(stub shim.ChaincodeStubInterface, args []string) peer.Response { //function, args := stub.GetFunctionAndParameters()

        // Condition checking
        if len(args) != 3 {
                return shim.Error("Incorrect number of arguments, required 3")
        }

        if args[1] == args[2] {
                return shim.Error("Invalid arguments, Trying to transfer in the same company")
        }
        // args[0]-empID | args[1]-from | args[2]-To
        employeeAsBytes, err := stub.GetState(args[0])

        if err != nil {
                return shim.Error(err.Error())
        }
        employee := Employee{}
        json.Unmarshal(employeeAsBytes, &employee)

        if employee.ParentCompany != "human" {
                return shim.Error("Only Human employees can be transferred")
        }

        if employee.EmployeeType != "contract" {
                return shim.Error("Only contract employees can be transferred")
        }

        if employee.CurrentCompany != args[1] {
                return shim.Error("This employee is not currently working in " + args[1] + " company")
        }

        employee.CurrentCompany = args[2]
        employeeAsBytes, _ = json.Marshal(employee)
        stub.PutState(args[0], employeeAsBytes)
        payload := []byte("Employee transfered Successfully")

        // Setting event for subscribers/listeners
        stub.SetEvent("sendEmployee", payload)
        return peer.Response{Status: 200, Message: "Record updated", Payload: payload}

}

/*
 * func name   : SetSalary
 * description : We use Private Data Collection concepts
 *               Setting Salary details of employee
 */

func (contract *ContractChaincode) SetSalary(stub shim.ChaincodeStubInterface, params []string) peer.Response {

        // Minimum of 2 args is needed - skipping the check for clarity
        // params[0]=Collection name
        // params[1]=Value for the token

        CollectionName := params[0]
        //TokenValue := params[1]
        EmpId := params[1]
        Salary := params[2]

        err := stub.PutPrivateData(CollectionName, EmpId, []byte(Salary))
        if err != nil {
                return shim.Error("Error=" + err.Error())
        }

        return shim.Success([]byte("true"))
}

/*
 * func name   : GetSalary
 * description : Getting Salary details of the employee
 *               depends on the PrivateDataCollection policy only authorized organization only can see the salary details
 */
func (contract *ContractChaincode) GetSalary(stub shim.ChaincodeStubInterface, params []string) peer.Response {
        // This is returned
        resultString := "{}"
        EmpId := params[0]

        // Read the open data
        dataOpen, err1 := stub.GetPrivateData("HumanHitachiContract", EmpId)
        if err1 != nil {
                return shim.Error("Error1=" + err1.Error())
        }

        // Read the acme private data
        dataSecret, err2 := stub.GetPrivateData("HumanPrivate", EmpId)

        accessError := "N.A."
        if err2 != nil {
                //return shim.Error("Error="+err1.Error())
                fmt.Println("Error2=" + err2.Error())
                accessError = err2.Error()
                dataSecret = []byte("**** Not Allowed ***")
        }

        // Returns the token value from 2 PDC + error
        resultString = "{Human Hitachi Salary:\"" + string(dataOpen) + "\", HumanSecretSalary:\"" + string(dataSecret) + "\" , error:\"" + accessError + "\"}"

        return shim.Success([]byte(resultString))

}

/*
 * func name   : getEmployeeTxHistory
 * description : Entry point of the function
 */

func (contract *ContractChaincode) getEmployeeTxHistory(stub shim.ChaincodeStubInterface, args []string) peer.Response {
        // Check the number of args
        if len(args) < 1 {
                return shim.Error("MUST provide EmployeeId !!!")
        }

        // Get the history for the employee
        historyQueryIterator, err := stub.GetHistoryForKey(args[0])

        // In case of error - return error
        if err != nil {
                return shim.Error("Error in fetching history !!!" + err.Error())
        }

        // Local variable to hold the history record
        var resultModification *queryresult.KeyModification
        counter := 0
        resultJSON := "["

        // Start a loop with check for more rows
        for historyQueryIterator.HasNext() {

                // Get the next record
                resultModification, err = historyQueryIterator.Next()

                if err != nil {
                        return shim.Error("Error in reading history record!!!" + err.Error())
                }

                // Append the data to local variable
                data := "{\"txn\":" + resultModification.GetTxId()
                data += " , \"value\": " + string(resultModification.GetValue()) + "}  "
                if counter > 0 {
                        data = ", " + data
                }
                resultJSON += data

                counter++
        }

        // Close the iterator
        historyQueryIterator.Close()

        // finalize the return string
        resultJSON += "]"
        resultJSON = "{ \"counter\": " + strconv.Itoa(counter) + ", \"txns\":" + resultJSON + "}"

        // return success
        return shim.Success([]byte(resultJSON))

}

/*
 * func name   : main
 * description : Entry point of the function
 */
func main() {
        fmt.Println("Started chain code")
        err := shim.Start(new(ContractChaincode))
        if err != nil {
                fmt.Println("Error Starting chain code : %s", err)
        }
}
