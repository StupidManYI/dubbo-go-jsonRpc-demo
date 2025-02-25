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
	"dubbo.apache.org/dubbo-go/v3"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"dubbo.apache.org/dubbo-go/v3/server"
	"github.com/dubbogo/gost/log/logger"
)

type GreetProvider struct {
}

func (*GreetProvider) SayHello(req string, req1 string, req2 string) (string, error) {
	return req + req1 + req2, nil
}

func main() {
	ins, err := dubbo.NewInstance(
		dubbo.WithProtocol(
			protocol.WithJSONRPC(),
			protocol.WithPort(20000)),
	)
	if err != nil {
		panic(err)
	}

	//JsonRpc
	srvJsonRpc, err := ins.NewServer()
	if err != nil {
		panic(err)
	}
	if err = srvJsonRpc.Register(&GreetProvider{}, nil, server.WithInterface("GreetProvider")); err != nil {
		panic(err)
	}
	if err := srvJsonRpc.Serve(); err != nil {
		logger.Error("srvJsonRpc err")
		logger.Error(err)
	}
}
