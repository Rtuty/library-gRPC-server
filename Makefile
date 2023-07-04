ifeq ($(OS),Windows_NT)
# SHELL=C:/Windows/System32/WindowsPowerShell/v1.0/powershell.exe
SHELL=C:/Program Files/Git/bin/bash.exe
else
SHELL=/usr/bin/bash
endif

SRC=C:/Users/kirill/Desktop

docker:
	#docker build -t library-service .
	#docker run --name=grpc-library-service -p 8000:8000 library-service
	docker-compose up --build library-service


prepare:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

generate_grpc_code: prepare
	protoc \
	--go_out="${SRC}/library-gRPC-server/internal/transport/library" \
 	--go_opt=paths=source_relative \
 	--go-grpc_out="${SRC}/library-gRPC-server/internal/transport/library" \
 	--go-grpc_opt=paths=source_relative library.proto


.PHONY: generate_grpc_code
.DEFAULT_GOAL=generate_grpc_code