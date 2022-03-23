package main

import (
	"fmt"

	"google.golang.org/grpc"
)

var (
	con *grpc.ClientConn // global variables are harmful
)

func connect() error {
	var err error
	// Set up a connection to the server
	addr := "localhost" + ":" + "8081"
	/*cr := insecure.NewCredentials()
	con, err = grpc.Dial(addr, grpc.WithTransportCredentials(cr))*/
	con, err = grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		return fmt.Errorf("rpc connect: %w", err)
	}

	return nil
}

func terminate() {
	if con != nil {
		_ = con.Close()
	}
}
