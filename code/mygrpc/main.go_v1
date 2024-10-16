package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "mygrpc/src/protos_ext"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Simuler une base de données en mémoire
	users := map[string]*pb.GetUserResponse{
		"1": {Id: "1", Name: "John Doe", Age: 30},
		"2": {Id: "2", Name: "Jane Doe", Age: 25},
	}

	user, exists := users[req.Id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func startServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startClient() {
	// Set up a connection to the server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.NewClient("passthrough:///localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	// Contact the server and print out its response.
	r, err := c.GetUser(ctx, &pb.GetUserRequest{Id: "1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetName())
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <server|client>", os.Args[0])
	}

	switch os.Args[1] {
	case "server":
		startServer()
	case "client":
		startClient()
	default:
		log.Fatalf("Usage: %s <server|client>", os.Args[0])
	}
}

// Pour lancer le serveur, exécutez la commande suivante:
// go run main.go server
// Pour lancer le client, exécutez la commande suivante:
// go run main.go client
// Vous devriez voir la sortie suivante:
// Date Time server listening at
// Date Time Greeting: John Doe
//
// Vous pouvez également tester le serveur avec un client gRPC généré en utilisant le fichier protos_ext.pb.go.
//
// Cela conclut le tutoriel sur la création d'un serveur gRPC simple en Go. Vous pouvez maintenant créer vos propres services gRPC pour communiquer entre les applications en utilisant RPC.
