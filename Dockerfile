FROM golang:1.20-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o library-service ./cmd/library-gRPC-server/main.go

CMD ["./library-service"]