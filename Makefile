generate_grpc_code:
	 C:\protobuf\bin\protoc.exe \
	--go_out=library \
 	--go_opt=paths=source_relative \
 	--go-grpc_out=library \
 	--go-grpc_opt=paths=source_relative library.proto