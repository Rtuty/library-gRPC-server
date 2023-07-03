package app

import (
	"database/sql"
	"google.golang.org/grpc"
	"log"
	"modules/internal/config"
	"modules/internal/db"
	"modules/internal/transport"
	"modules/internal/transport/library"
	"net"
)

func RunServer() {
	cf, err := config.GetConnection()
	if err != nil {
		panic(err)
	}

	con, err := sql.Open("mysql", cf.User+":"+cf.Passwd+"@/"+cf.Dbname)
	if err != nil {
		panic(err)
	}

	strg := db.NewMySqlRepository(con)

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegister := grpc.NewServer()
	service := &transport.LibraryServer{Db: strg}

	library.RegisterLibraryServer(serverRegister, service)

	if err = serverRegister.Serve(lis); err != nil {
		log.Fatalf("cannot serve server")
	}
}
