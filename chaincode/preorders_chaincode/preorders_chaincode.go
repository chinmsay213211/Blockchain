package main

// This is sample lot_chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"bytes"
	"github.com/lib"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Files describes base64 file's state
type Files struct {
	Name   string `json:"Name"`
	Binary string `json:"Binary"`
	JSON   string `json:"JSON"`
}

// PreOrder describes chaincode's state
type PreOrder struct {
	Hash              string      `json:"Hash"`
	ClientUserHashes  []string    `json:"ClientUserHashes"`
	StatusesHashes    []Statuses  `json:"StatusesHashes"`
	CompanyHash       string      `json:"CompanyHash"`
	ManagerUserHashes []string    `json:"ManagerUserHashes"`
    DispatchNo        string      `json:"DispatchNo"`
    ShipperHash           string      `json:"ShipperHash"`
    ShipToHash            string      `json:"ShipToHash"`
    Deleted           bool        `json:"Deleted"`
    Active            bool        `json:"Active"`
    ExpDate           int         `json:"ExpDate"`
    Timestamp         int         `json:"Timestamp"`
}

// Statuses defined separately for better readability
type Statuses struct {
    StatusType string `json:"type"`
    UserIdHash string `json:"UserHash"`
    StatusHash string `json:"StatusHash"`
}


// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke chaincode methods
// ===========================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "create" {
		return t.create(stub, args)
	} else if function == "update" {
		return t.update(stub, args)
	} else if function == "get" {
		return t.get(stub, args)
	} else if function == "query" {
		return t.query(stub, args[0])
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// Create method creates entity
func (t *SimpleChaincode) create(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println(fmt.Sprintf("create executed with args: %+v", args))

	var err error
	if len(args) != 1 {
		str := fmt.Sprintf("Incorrect number of arguments %d.", len(args))
		fmt.Println(str)
		return shim.Error(str)
	}

	if len(args[0]) <= 0 {
		str := fmt.Sprintf("JSON must be non-empty string %+v", args)
		fmt.Println(str)
		return shim.Error(str)
	}

	data := &PreOrder{}
	err = json.Unmarshal([]byte(args[0]), data)
	if err != nil {
		str := fmt.Sprintf("Failed to parse JSON: %+v", err)
		fmt.Println(str)
		return shim.Error(str)
	}

	dataAsBytes, err := stub.GetState(data.Hash)
	if err != nil {
		str := fmt.Sprintf("Failed to get: %s", data.Hash)
		fmt.Println(str)
		return shim.Error(str)
	} else if dataAsBytes != nil {
		str := fmt.Sprintf("Record already exists: %s", data.Hash)
		fmt.Println(str)
		return shim.Error(str)
	}
	dataJSONasBytes, err := json.Marshal(data)
	if err != nil {
		str := fmt.Sprintf("Could not marshal %+v", err.Error())
		fmt.Println(str)
		return shim.Error(str)
	}


	if !lib.Validate("companies_chaincode", data.CompanyHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "companies_chaincode", data.CompanyHash)
		return shim.Error(errStr)
	}

	if !lib.Validate("locations_chaincode", data.ShipperHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "locations_chaincode", data.ShipperHash)
		return shim.Error(errStr)
	}

	if !lib.Validate("locations_chaincode", data.ShipToHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "locations_chaincode", data.ShipToHash)
		return shim.Error(errStr)
	}


	for i := 0; i < len(data.ClientUserHashes); i++ {
		if !lib.Validate("users_chaincode", data.ClientUserHashes[i], stub) {
			errStr := fmt.Sprintf("Failed to validate %s on %s", "users_chaincode", data.ClientUserHashes[i])
			return shim.Error(errStr)
		}
	}

	for i :=0; i < len(data.ManagerUserHashes); i++ {
		if !lib.Validate("users_chaincode", data.ManagerUserHashes[i], stub) {
			errStr := fmt.Sprintf("Failed to validate %s on %s", "users_chaincode", data.ManagerUserHashes[i])
			return shim.Error(errStr)
		}
	}

	err = stub.PutState(data.Hash, dataJSONasBytes)
	if err != nil {
		str := fmt.Sprintf("Could not put state %+v", err.Error())
		fmt.Println(str)
		return shim.Error(str)
	}

	fmt.Println("Sucessfully executed");
	return shim.Success(nil)
}

func (t *SimpleChaincode) update(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 1 {
		str := fmt.Sprintf("Incorrect number of arguments %d.", len(args))
		fmt.Println(str)
		return shim.Error(str)
	}

	if len(args[0]) <= 0 {
		str := fmt.Sprintf("JSON must be non-empty string %+v", args)
		fmt.Println(str)
		return shim.Error(str)
	}

	data := &PreOrder{}
	err = json.Unmarshal([]byte(args[0]), data)

	dataAsBytes, err := stub.GetState(data.Hash)

	if err != nil {
		str := fmt.Sprintf("Failed to get: %+v", err.Error());
		fmt.Println(str)
		return shim.Error(str)
	} else if dataAsBytes == nil {
		str := fmt.Sprintf("Record does not exists: %s", data.Hash)
		fmt.Println(str)
		return shim.Error(str)
	}

	dataJSONasBytes, err := json.Marshal(data)

	if err != nil {
		str := fmt.Sprintf("Can not marshal %+v", err.Error())
		fmt.Println(str)
		return shim.Error(str)
	}

	if !lib.Validate("companies_chaincode", data.CompanyHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "companies_chaincode", data.CompanyHash)
		return shim.Error(errStr)
	}

	for i := 0; i < len(data.ClientUserHashes); i++ {
		if !lib.Validate("users_chaincode", data.ClientUserHashes[i], stub) {
			errStr := fmt.Sprintf("Failed to validate %s on %s", "users_chaincode", data.ClientUserHashes[i])
			return shim.Error(errStr)
		}
	}

	for i :=0; i < len(data.ManagerUserHashes); i++ {
		if !lib.Validate("users_chaincode", data.ManagerUserHashes[i], stub) {
			errStr := fmt.Sprintf("Failed to validate %s on %s", "users_chaincode", data.ManagerUserHashes[i])
			return shim.Error(errStr)
		}
	}

	err = stub.PutState(data.Hash, dataJSONasBytes)

	if err != nil {
		str := fmt.Sprintf("Can not put state %+v", err.Error())
		fmt.Println(str)
		return shim.Error(str)
	}

	fmt.Println("Successfully updated")
	return shim.Success(nil)
}

func (t *SimpleChaincode) get(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		str := fmt.Sprintf("Incorrect number(%d) of arguments", len(args))
		fmt.Println(str)
		return shim.Error(str)
	}

	hash := args[0]
	valAsBytes, err := stub.GetState(hash)

	if err != nil {
		str := fmt.Sprintf("Failed to get state %+v", err.Error())
		fmt.Println(str)
		return shim.Error(str)
	} else if valAsBytes == nil {
		str := fmt.Sprintf("Record does not exist %s", hash)
		fmt.Println(str)
		return shim.Error(str)
	}

	fmt.Println("Successfully got")
	return shim.Success(valAsBytes)
}

func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, queryString string) pb.Response {
	fmt.Println(fmt.Sprintf("query started %s", queryString))
	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		fmt.Println(fmt.Sprintf("incorrect query: %s", queryString))
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
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

	fmt.Println("Sucessfully queried")
	return shim.Success(buffer.Bytes())
}
