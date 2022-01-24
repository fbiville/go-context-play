package server

import (
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

func SomeServer() (net.Listener, error) {
	listener, err := net.Listen("tcp", "localhost:")
	if err != nil {
		return nil, err
	}
	go http.Serve(listener, waitAndServe204())
	return listener, nil
}

func waitAndServe204() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			handleError(resp, err)
			return
		}
		timeoutInSecs, err := strconv.Atoi(string(body))
		if err != nil {
			handleError(resp, err)
			return
		}
		time.Sleep(time.Second * time.Duration(timeoutInSecs))
		resp.WriteHeader(204)
		resp.Write([]byte{})
	}
}

func handleError(resp http.ResponseWriter, err error) {
	resp.WriteHeader(500)
	resp.Header().Add("Content-Type", "text/plain")
	resp.Write([]byte(err.Error()))
}
