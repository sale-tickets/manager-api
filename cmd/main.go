package main

import (
	"fmt"

	"github.com/sale-tickets/manager-api/internal/common/connection"
	"github.com/sale-tickets/manager-api/internal/router"
)

func main() {
	connection.Init()

	chanStartedGrpc := make(chan bool, 1)
	var chanError = make(chan error, 1)

	go func() {
		router.GrpcServer(chanStartedGrpc, chanError)
	}()

	go func() {
		router.HttpServer(chanStartedGrpc, chanError)
	}()

	err := <-chanError
	if err != nil {
		fmt.Println("error run main: ", err.Error())
		return
	}
}
