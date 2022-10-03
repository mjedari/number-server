# Number Server Instructions

## Docker
#### To build and run
```
docker build --tag number-server .
docker run -it -p 4000:4000 number-server
```

## Manually
#### To run
``
go run *.go
``

### To build and run
``
go build -o number-server *.go && ./number-server
``


## Connect client
``
telnet localhot 4000 
``