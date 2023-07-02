generate_grpc_code:
	 C:\protobuf\bin\protoc.exe \
	--go_out=C:\Users\kirill\Desktop\library-gRPC-server\internal\transport\library \
 	--go_opt=paths=source_relative \
 	--go-grpc_out=C:\Users\kirill\Desktop\library-gRPC-server\internal\transport\library \
 	--go-grpc_opt=paths=source_relative library.proto