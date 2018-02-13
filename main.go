package main

import (
	"fmt"

	socks5 "github.com/armon/go-socks5"
)

func main() {
	// Create a SOCKS5 server
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 2222
	addr := "0.0.0.0:2222"
	fmt.Println("Starting server on", addr)
	if err := server.ListenAndServe("tcp", addr); err != nil {
		panic(err)
	}
}
