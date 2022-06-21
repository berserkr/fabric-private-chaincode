package chaincode

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type SimpleAsset struct {
}

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	switch f, _ := stub.GetFunctionAndParameters(); f {
	case "store":
		return storeAsset(stub)
	case "retrieve":
		return retrieveAsset(stub)
	}

	return shim.Error("unknown function")
}

func storeAsset(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()

	if len(args) < 2 {
		return shim.Error("not enough arguments")
	}

	assetName, value := args[0], args[1]

	if err := stub.PutState(assetName, []byte(value)); err != nil {
		return shim.Error("something went wrong")
	}

	return shim.Success([]byte("OK"))
}

func retrieveAsset(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()

	if len(args) < 1 {
		return shim.Error("not enough arguments")
	}

	assetName := args[0]

	value, err := stub.GetState(assetName)
	if err != nil {
		return shim.Error("something went wrong")
	}

	if len(value) == 0 {
		shim.Success([]byte("NOT FOUND"))
	}

	return shim.Success([]byte(fmt.Sprintf("%s:%s", assetName, value)))
}
