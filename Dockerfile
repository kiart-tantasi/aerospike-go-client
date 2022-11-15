FROM golang:1.19

WORKDIR /app

COPY main.go ./

RUN go mod init aerospike-go-client
RUN go get github.com/aerospike/aerospike-client-go

RUN go build -o app main.go

CMD [ "./app" ]
