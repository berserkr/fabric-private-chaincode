/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package cbdc

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hyperledger-labs/fabric-smart-client/integration"
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/common"
	"github.com/hyperledger/fabric-private-chaincode/samples/demos/cbdc/views/cbdc_users"
	"github.com/stretchr/testify/assert"
)

func TestFlow(t *testing.T) {

	// setup fabric network
	ii, err := integration.Generate(23000, false, Topology()...)
	assert.NoError(t, err)
	ii.Start()
	defer ii.Stop()

	// data provider flow
	_, err = ii.Client("alice").CallView("Register", common.JSONMarshall(&cbdc_users.Register{
		UserID:     "01",
		UserVK: []byte("abc"),
		AttestationData: []byte("data"),
		AttestationDataSig:   []byte("sig"),
	}))
	assert.NoError(t, err)


}
