package chaincode

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type AttestationRegistry struct {
}

func (t *AttestationRegistry) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *AttestationRegistry) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	switch f, _ := stub.GetFunctionAndParameters(); f {
	case "store":
		return storeAttestation(stub)
	case "retrieve":
		return retrieveAttestation(stub)
	}

	return shim.Error("unknown function")
}

func storeAttestation(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()

	if len(args) < 2 {
		return shim.Error("not enough arguments")
	}

	attestationHash, value := args[0], args[1]

	if msg, err := verifyAttestation(value); err != nil {
		return shim.Error(msg)
	}

	if err := stub.PutState(attestationHash, []byte(value)); err != nil {
		return shim.Error("something went wrong")
	}

	return shim.Success([]byte("OK"))
}

func retrieveAttestation(stub shim.ChaincodeStubInterface) pb.Response {
	_, args := stub.GetFunctionAndParameters()

	if len(args) < 1 {
		return shim.Error("not enough arguments")
	}

	attestationHash := args[0]

	value, err := stub.GetState(attestationHash)
	if err != nil {
		return shim.Error("something went wrong")
	}

	if len(value) == 0 {
		shim.Success([]byte("NOT FOUND"))
	}

	return shim.Success([]byte(fmt.Sprintf("%s:%s", attestationHash, value)))
}
