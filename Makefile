build:
	go build -o server

release: build
	docker build -t tlake/go-socks5-proxy:latest .
	docker push tlake/go-socks5-proxy:latest

.PHONY:
	build release
