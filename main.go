package main

import (
	"context"
	"fmt"
	"github.com/fbiville/context-play/pkg/client"
	"github.com/fbiville/context-play/pkg/server"
	"net/http"
	"time"
)

func main() {
	aServer, err := server.SomeServer()
	if err != nil {
		panic(err)
	}
	defer aServer.Close()

	aClient := client.SomeClient(&http.Client{}, aServer.Addr())
	now := time.Now()
	defer func() {
		fmt.Printf("Time elapsed is: %v second(s)", time.Now().Sub(now).Seconds())
	}()

	timeout, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()

	err = aClient.SomeComplexCalls(timeout)
	fmt.Printf("%v\n", err)
}
