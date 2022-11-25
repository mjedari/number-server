test:
	go test ` go list ./... | grep -v vendor`

build:
	go build -o number-server *.go && ./number-server

build-by-docker:
	docker build --tag number-server .