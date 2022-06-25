package chaincode

import (
	"fmt"
)


/*
func add (a int, b int) int {
    return a + b
}
*/
func verifyAttestation(attestation string) (string, error) {

	fmt.Println("> verify " + attestation)

	return "OK", nil;
}



