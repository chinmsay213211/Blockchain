package main

// This is sample deliveryVoucher_chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/lib"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Files with ContentType
type Drops struct {
	FromLocationHash string `json:"FromLocationHash"`
	ToLocationHash   string `json:"ToLocationHash"`
	LotHash          string `json:"LotHash"`
}
// DeliveryVoucher describes chaincode's structure
type DeliveryVoucher struct {
	Hash                       string `json:"Hash"`
	Number                     int `json:"Number"`
	DeliveryToCompanyHash      string `json:"DeliveryToCompanyHash"`
	ManagerUserHash            string `json:"ManagerUserHash"`
	OrderTransportationHash    string `json:"OrderTransportationHash"`
	Drops                      Drops  `json:"Drops,omitempty"`
	DriverUserHash             string `json:"DriverUserHash"`
	TruckHash                  string `json:"TruckHash"`
	OpeningKm                  int    `json:"OpeningKm"`
	ClosingKm                  int    `json:"ClosingKm"`
	TotalKm                    int    `json:"TotalKm"`
	TimeStart                  int    `json:"TimeStart"`
	TimeEnd                    int    `json:"TimeEnd"`
	WorkingHours               int    `json:"WorkingHours"`
	DrivingHours               int    `json:"DrivingHours"`
	RestHours                  int    `json:"RestHours"`
	PdnHours                   int    `json:"PdnHours"`
	DeliveryStatus             string `json:"DeliveryStatus"`
	ConfirmedByManagerUserHash string `json:"ConfirmedByManagerUserHash"`
	ClientSignatureImage       string `json:"ClientSignatureImage"`
	ClientSignedDvPhoto        string `json:"ClientSignedDvPhoto"`
	ClientBadgePhoto           string `json:"ClientBadgePhoto"`
	Timestamp                  int    `json:"Timestamp"`
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
	fmt.Println("invoke is running " + function)

	if function == "create" {
		return t.create(stub, args)
	} else if function == "update" {
		return t.update(stub, args)
	} else if function == "get" {
		return t.get(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// Create method creates entity
func (t *SimpleChaincode) create(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments")
	}

	if len(args[0]) <= 0 {
		return shim.Error("JSON must be non-empty string")
	}

	data := &DeliveryVoucher{}
	err = json.Unmarshal([]byte(args[0]), data)

	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to parse JSON: %+v", err))
	}

	dataAsBytes, err := stub.GetState(data.Hash)

	if err != nil {
		return shim.Error("Failed to get: " + data.Hash)
	} else if dataAsBytes != nil {
		return shim.Error("Record already exists: " + data.Hash)
	}

	dataJSONasBytes, err := json.Marshal(data)

	if err != nil {
		return shim.Error(err.Error())
	}

	if !lib.Validate("companies_chaincode", data.DeliveryToCompanyHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "companies_chaincode", data.DeliveryToCompanyHash)
		return shim.Error(errStr)
	}

	if !lib.Validate("orders_transportation_chaincode", data.OrderTransportationHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "orders_transportation_chaincode", data.OrderTransportationHash)
		return shim.Error(errStr)
	}

	if !lib.Validate("users_chaincode", data.DriverUserHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "users_chaincode", data.DriverUserHash)
		return shim.Error(errStr)
	}

	if !lib.Validate("trucks_chaincode", data.TruckHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "trucks_chaincode", data.TruckHash)
		return shim.Error(errStr)
	}

	if !lib.Validate("users_chaincode", data.ConfirmedByManagerUserHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "users_chaincode", data.ConfirmedByManagerUserHash)
		return shim.Error(errStr)
	}

	err = stub.PutState(data.Hash, dataJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) update(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments")
	}

	if len(args[0]) <= 0 {
		return shim.Error("JSON must be non-empty string")
	}

	data := &DeliveryVoucher{}
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
		return shim.Error(err.Error())
	}

	if !lib.Validate("companies_chaincode", data.DeliveryToCompanyHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "companies_chaincode", data.DeliveryToCompanyHash)
		return shim.Error(errStr)
	}

	if !lib.Validate("orders_transportation_chaincode", data.OrderTransportationHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "orders_transportation_chaincode", data.OrderTransportationHash)
		return shim.Error(errStr)
	}

	if !lib.Validate("users_chaincode", data.DriverUserHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "users_chaincode", data.DriverUserHash)
		return shim.Error(errStr)
	}

	if !lib.Validate("trucks_chaincode", data.TruckHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "trucks_chaincode", data.TruckHash)
		return shim.Error(errStr)
	}

	if !lib.Validate("users_chaincode", data.ConfirmedByManagerUserHash, stub) {
		errStr := fmt.Sprintf("Failed to validate %s on %s", "users_chaincode", data.ConfirmedByManagerUserHash)
		return shim.Error(errStr)
	}

	err = stub.PutState(data.Hash, dataJSONasBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) get(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error(fmt.Sprintf("Incorrect number(%d) of arguments", len(args)))
	}

	hash := args[0]
	valAsBytes, err := stub.GetState(hash)

	if err != nil {
		jsonResp := "\"Error\" : \"Failed to get state for " + hash + "\"}"
		return shim.Error(jsonResp)
	} else if valAsBytes == nil {
		jsonResp := "\"Error\" : Record does not exist: " + hash + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsBytes)
}