/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/dubbogo/gost/log/logger"
)

func main() {
	ins, err := dubbo.NewInstance(
		dubbo.WithName("dubbo_multirpc_client"),
	)
	if err != nil {
		panic(err)
	}
	//JsonRpc
	cliJsonRpc, err := ins.NewClient(
		client.WithClientURL("127.0.0.1:20000"),
		client.WithClientProtocolJsonRPC(),
		client.WithClientSerialization(constant.JSONSerialization),
	)
	if err != nil {
		panic(err)
	}
	connJsonRpc, err := cliJsonRpc.Dial("GreetProvider")
	if err != nil {
		panic(err)
	}
	var respJsonRpc string
	if err := connJsonRpc.CallUnary(context.Background(), []interface{}{"hello", "new", "jsonrpc"}, &respJsonRpc, "SayHello"); err != nil {
		logger.Errorf("GreetProvider.Greet err: %s", err)
		return
	}
	logger.Infof("Get jsonrpc Response: %s", respJsonRpc)

}
