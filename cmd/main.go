package main 

import(
	proto "github.com/Dev79844/chat-grpc/gen"
	"google.golang.org/grpc"
	server "github.com/Dev79844/chat-grpc/server"
	"net"
	"log"
	"fmt"
)

func main(){
	grpcServer := grpc.NewServer()

	var conn []*server.Connection

	pool := &server.Pool{
		Connection: conn,
	}

	proto.RegisterBroadcastServer(grpcServer, pool)

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Error creating the server %v", err)
	}

	fmt.Println("Server started at port 3000")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error creating the server %v", err)
	}
}
