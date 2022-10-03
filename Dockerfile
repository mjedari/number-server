FROM golang:1.19.1

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN go build -o number-server

CMD [ "./number-server" ]