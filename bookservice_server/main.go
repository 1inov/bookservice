package main

import (
	"bookservice/server"
	"bookservice/server/grpc_server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// создание grpc сервера
	s := grpc.NewServer()
	// создание структуры для grpc сервера
	srv := &grpc_server.GRPCServer{}
	// регистрация структуры на интерфейс grpc сервера
	server.RegisterBaseServer(s, srv)

	// создание слушателя
	l, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatalf("Service can't create listener : %s", err)
	}

	// Подключение слушателя к grpc серверу
	if err := s.Serve(l); err != nil {
		log.Fatalf("Service can't serve grpcserver : %s", err)
	}
}
