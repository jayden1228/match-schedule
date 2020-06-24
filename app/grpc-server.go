package app

import (
	"log"
	"match-schedule/app/handlers"
	"match-schedule/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunServer with port
func RunServer(port string) {
	s := grpc.NewServer()
	proto.RegisterMatchServiceServer(s, &handlers.GrpcServer{})
	reflection.Register(s)

	go func() {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatal("Failed to listen: ", err)
		}
		defer lis.Close()

		log.Println("Start Grpc Server ", port)
		if err = s.Serve(lis); err != nil {
			log.Fatal("Failed to serve: ", err)
		}
	}()
}
