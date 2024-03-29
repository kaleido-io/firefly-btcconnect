// Copyright © 2022 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bitcoin

import (
	"context"
	"testing"

	"github.com/btcsuite/btcwallet/netparams"
	"github.com/hyperledger/firefly-btcconnect/internal/jsonrpc"
	"github.com/hyperledger/firefly-btcconnect/mocks/jsonrpcmocks"
	"github.com/hyperledger/firefly-common/pkg/config"
	"github.com/hyperledger/firefly-common/pkg/ffresty"
	"github.com/stretchr/testify/assert"
)

func newTestConnector(t *testing.T) (*btcConnector, *jsonrpc.RPCClient) {

	wsclient := &jsonrpcmocks.WsClient{}
	mRPC := jsonrpc.NewRPCClient(context.Background(), wsclient, (&netparams.TestNet3Params).Params)
	config.RootConfigReset()
	conf := config.RootSection("unittest")
	InitConfig(conf)
	conf.Set("auth.username", "user1")
	conf.Set("auth.password", "pwd")
	c := &btcConnector{}
	conf.Set(ffresty.HTTPConfigURL, "http://backend.example.invalid")
	err := c.Init(context.Background(), conf)
	assert.NoError(t, err)
	c.backend = mRPC
	return c, mRPC
}
