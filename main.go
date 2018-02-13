package main

import (
	"fmt"
	"os"

	socks5 "github.com/armon/go-socks5"
)

func main() {
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PROXY_PORT")
	if port == "" {
		port = "2222"
	}

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Println("Starting server on", addr)
	if err := server.ListenAndServe("tcp", addr); err != nil {
		panic(err)
	}
}
