/*
   Copyright IBM Corp. All Rights Reserved.

   SPDX-License-Identifier: Apache-2.0
*/

package cbdc_users

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger-labs/fabric-smart-client/platform/fabric/services/fpc"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/assert"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"

	"github.com/pkg/errors"
)

type Register struct {
	UserID             string
	UserVK             []byte
	AttestationData    []byte
	AttestationDataSig []byte
}

type RegisterView struct {
	*Register
}

func (c *RegisterView) Call(context view.Context) (interface{}, error) {
	fmt.Printf("Register new user")

	cid := "registry"
	f := "store"

	v, err := fpc.GetDefaultChannel(context).EnclaveRegistry().IsAvailable()
	assert.NoError(err, "failed checking availability of the enclave registry")
	assert.True(v, "the enclave registry is not available")

	v, err = fpc.GetDefaultChannel(context).EnclaveRegistry().IsPrivate(cid)
	assert.NoError(err, "failed checking echo deployment")
	assert.True(v, "echo should be an FPC")

	var arg []string
	arg = append(arg, c.UserID)
	arg = append(arg, string(c.AttestationData))

	if _, err := fpc.GetDefaultChannel(context).Chaincode(cid).Invoke(f, arg).Call(); err != nil {
		return nil, errors.Wrapf(err, "error invoking %s", f)
	}

	fmt.Println("Patient data successfully registered! thanks")

	return nil, nil
}

type RegisterViewFactory struct{}

func (c *RegisterViewFactory) NewView(in []byte) (view.View, error) {
	f := &RegisterView{Register: &Register{}}
	if err := json.Unmarshal(in, f.Register); err != nil {
		return nil, err
	}
	return f, nil
}
