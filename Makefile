run:
	@go run main.go

build:
	@go build -o $(app_name) main.go

start: clean build run

build-by-docker:
	@docker build --tag $(app_name) .

run-by-docker:
	@docker run -it -p 4000:4000 --network host $(app_name)

start-by-docker: build-by-docker run-by-docker

clean:
	@rm -f $(app_name)
	@echo "project cleaned!"

test:
	@echo "Start to test..."
	@go test ` go list ./... | grep -v vendor`

app_name:= "number-server"