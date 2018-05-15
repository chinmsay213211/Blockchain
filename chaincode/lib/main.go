package lib

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
)

func Validate(chaincodeID string, hash string, stub shim.ChaincodeStubInterface) bool {
	//return true
	args := [2]string {"get", hash}
	for i := 0; i < 3; i++ {
		invokeArgs := make([][]byte, len(args))
		for i, arg := range args {
			invokeArgs[i] = []byte(arg)
		}

		response := stub.InvokeChaincode(chaincodeID, invokeArgs, "mainchannel")

		if response.Status == shim.OK {
			fmt.Println(fmt.Sprintf("Sucessfully validated %s on %s", chaincodeID, hash))
			return true;
		}

		fmt.Println(fmt.Sprintf("Failed to validate %s on %s", chaincodeID, hash))
	}
	return false;
}