package client

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

type api struct {
	client        *http.Client
	serverAddress string
}

func SomeClient(client *http.Client, address net.Addr) *api {
	return &api{
		client:        client,
		serverAddress: fmt.Sprintf("http://%s", address.String()),
	}
}

func (api *api) SomeSlowOperations(ctx context.Context) error {
	if err := api.slowOperation(ctx, 10*time.Second); err != nil {
		return err
	}
	if err := api.slowOperation(ctx, 10*time.Second); err != nil {
		return err
	}
	return nil
}

func (api *api) slowOperation(ctx context.Context, wait time.Duration) error {
	waitRequest := strings.NewReader(fmt.Sprintf("%d", int(wait.Seconds())))
	request, err := http.NewRequestWithContext(ctx, "GET", api.serverAddress, waitRequest)
	if err != nil {
		panic(err)
	}
	resp, err := api.client.Do(request)
	if err != nil {
		fmt.Println("Error occurred")
		return err
	}
	fmt.Printf("Response: %s\n", resp.Status)
	return nil
}
