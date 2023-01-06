package engine

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

func PrepareGrpcServer() (*grpc.Server, net.Listener) {
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")

	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return grpc.NewServer(), lis
}

func RunGrpcServer(s *grpc.Server, lis net.Listener) {
	log.Printf("cmd listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
