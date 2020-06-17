package app

import (
	"match-schedule/app/handlers"
	"match-schedule/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunServer with port
func RunServer(port string) {
	s := grpc.NewServer()
	proto.RegisterDemoServiceServer(s, &handlers.GrpcServer{})
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
