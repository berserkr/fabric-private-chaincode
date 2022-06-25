package main

import (
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	fpc "github.com/hyperledger/fabric-private-chaincode/ecc_go/chaincode"
	"github.com/hyperledger/fabric-private-chaincode/samples/demos/cbdc/registry/chaincode"
)

func main() {
	ccid := os.Getenv("CHAINCODE_PKG_ID")
	addr := os.Getenv("CHAINCODE_SERVER_ADDRESS")

	// create private chaincode
	privateChaincode := fpc.NewPrivateChaincode(&chaincode.AttestationRegistry{})

	// start chaincode as a service
	server := &shim.ChaincodeServer{
		CCID:    ccid,
		Address: addr,
		CC:      privateChaincode,
		TLSProps: shim.TLSProperties{
			Disabled: true, // just for testing good enough
		},
	}

	if err := server.Start(); err != nil {
		panic(err)
	}
}