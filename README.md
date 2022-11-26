# Number Server Instructions

### Quick start

1. run server
   ```bash
   # by docker
   make start-by-docker
   ```
   OR
   ```bash
   # manually
   make start
   ```

2. connect by telnet:
    ```bash
    tellnet localhost 4000
    ```

----------

### Manually detailed

#### Run
```bash
go run main.go
```
#### Build and run
```bash
go build -o number-server *.go && ./number-server
```

#### To build and run by docker
```bash
docker build --tag number-server .
docker run -it -p 4000:4000 --network host number-server
```

### Client connection
The program listens on port `4000` on telnet IPC. you can communicate only in this way
```bash
telnet localhost 4000 
```

### Testing
To run tests just run this command:
```bash
make test
```