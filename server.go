package chat111

import (
	chatserv "github.com/chat111"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	streamPort = "8012"
)

func main() {
	listen, err := net.Listen("tcp", ":"+streamPort)
	if err != nil {
		log.Fatalf("Can't launch at %v :: err:%v", streamPort, err)
	}
	log.Println("Launched at: " + streamPort)
}

grpcserver := grpc.NewServer()

cs := chatserver.ChatServer{}
chatserver.RegisterServicesServer(grpcserver, &cs)

err = grpcserver.Serve(listen)
if err != nil {
	log.Fatalf("Failed to start gRPC Server err: " + err)
}