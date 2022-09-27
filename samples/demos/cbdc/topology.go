/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package cbdc

import (
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/api"
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/fabric"
	"github.com/hyperledger-labs/fabric-smart-client/integration/nwo/fsc"
	cbdc_users "github.com/hyperledger/fabric-private-chaincode/samples/demos/cbdc/views/cbdc_users"
)

func Topology() []api.Topology {
	fabricTopology := fabric.NewDefaultTopology()
	fabricTopology.AddOrganizationsByName("Org1", "Org2", "Org3")
	fabricTopology.AddFPC("registry", "fpc/fpc-registry")
	fscTopology := fsc.NewTopology()

	// data provider
	aliceNode := fscTopology.AddNodeByName("alice")
	aliceNode.AddOptions(fabric.WithOrganization("Org1"))
	aliceNode.RegisterViewFactory("Register", &cbdc_users.RegisterViewFactory{})

	return []api.Topology{fabricTopology, fscTopology}
}
