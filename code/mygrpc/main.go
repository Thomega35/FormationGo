package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	pb "mygrpc/src/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users map[string]*pb.GetUserResponse
}

func (s *server) GetUserRpc(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.GetUser(ctx, req)
}

func (s *server) ListUsersRpc(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return s.ListUsers(ctx, req)
}

func (s *server) CreateUserRpc(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.CreateUser(ctx, req)
}

func (s *server) DeleteUserRpc(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.DeleteUser(ctx, req)
}

func (s *server) GetUserHttp(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return s.GetUser(ctx, req)
}

func (s *server) ListUsersHttp(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return s.ListUsers(ctx, req)
}

func (s *server) CreateUserHttp(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return s.CreateUser(ctx, req)
}

func (s *server) DeleteUserHttp(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return s.DeleteUser(ctx, req)
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, exists := s.users[req.Id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	var users []*pb.GetUserResponse
	for _, user := range s.users {
		users = append(users, user)
	}
	return &pb.ListUsersResponse{Users: users}, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	id := fmt.Sprintf("%d", len(s.users)+1)
	user := &pb.GetUserResponse{Id: id, Name: req.Name, Age: req.Age}
	s.users[id] = user
	return &pb.CreateUserResponse{Id: id}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, exists := s.users[req.Id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	delete(s.users, req.Id)
	return &pb.DeleteUserResponse{Id: req.Id}, nil
}

func startGrpcServer() *grpc.Server {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	service := &server{users: make(map[string]*pb.GetUserResponse)}
	pb.RegisterUserServiceServer(s, service)

	log.Printf("gRPC server listening at %v", lis.Addr())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return s
}

func startHttpServer() *http.Server {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not register service handler: %v", err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Println("HTTP server listening at :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return srv
}

func callRpcClient(c pb.UserServiceClient, ctx context.Context) {
	// Example: CreateUser
	start := time.Now()
	createResp, err := c.CreateUserRpc(ctx, &pb.CreateUserRequest{Name: "Alice", Age: 22})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	duration := time.Since(start)
	log.Printf("RPC CreateUser took %v", duration)
	log.Printf("Created User ID: %s", createResp.GetId())

	// Example: ListUsers
	start = time.Now()
	listResp, err := c.ListUsersRpc(ctx, &pb.ListUsersRequest{})
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
	duration = time.Since(start)
	log.Printf("RPC ListUsers took %v", duration)
	for _, user := range listResp.GetUsers() {
		log.Printf("User: %s, Name: %s, Age: %d", user.GetId(), user.GetName(), user.GetAge())
	}

	// Example: DeleteUser
	start = time.Now()
	deleteResp, err := c.DeleteUserRpc(ctx, &pb.DeleteUserRequest{Id: createResp.GetId()})
	if err != nil {
		log.Fatalf("could not delete user: %v", err)
	}
	duration = time.Since(start)
	log.Printf("RPC DeleteUser took %v", duration)
	log.Printf("Deleted User ID: %s", deleteResp.GetId())
}

func callHttpClient() {
	client := &http.Client{}

	// Example: CreateUser
	start := time.Now()
	createReq := &pb.CreateUserRequest{Name: "Alice", Age: 22}
	reqBody, _ := json.Marshal(createReq)
	resp, err := client.Post("http://localhost:8080/v1/createuser", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	defer resp.Body.Close()
	var createResp pb.CreateUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
		log.Fatalf("could not decode response: %v", err)
	}
	duration := time.Since(start)
	log.Printf("HTTP CreateUser took %v", duration)
	log.Printf("Created User ID: %s", createResp.GetId())

	// Example: ListUsers
	start = time.Now()
	resp, err = client.Get("http://localhost:8080/v1/listusers")
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
	defer resp.Body.Close()
	var listResp pb.ListUsersResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		log.Fatalf("could not decode response: %v", err)
	}
	duration = time.Since(start)
	log.Printf("HTTP ListUsers took %v", duration)
	for _, user := range listResp.GetUsers() {
		log.Printf("User: %s, Name: %s, Age: %d", user.GetId(), user.GetName(), user.GetAge())
	}

	// Example: DeleteUser
	start = time.Now()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/v1/deleteuser/%s", createResp.GetId()), nil)
	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}
	resp, err = client.Do(req)
	if err != nil {
		log.Fatalf("could not delete user: %v", err)
	}
	defer resp.Body.Close()
	var deleteResp pb.DeleteUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&deleteResp); err != nil {
		log.Fatalf("could not decode response: %v", err)
	}
	duration = time.Since(start)
	log.Printf("HTTP DeleteUser took %v", duration)
	log.Printf("Deleted User ID: %s", deleteResp.GetId())
}

func startClient() {
	// Set up a connection to the gRPC server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	// Call RPC functions
	fmt.Println("--- RPC ---")
	callRpcClient(c, ctx)

	// Call HTTP functions
	fmt.Println("--- HTTP ---")
	callHttpClient()
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <server|client|all>", os.Args[0])
	}

	switch os.Args[1] {
	case "server":
		startGrpcServer()
		startHttpServer()
		// Block main goroutine so servers continue running
		select {}
	case "client":
		startClient()
	case "all":
		done := make(chan struct{})
		go func() {
			grpcServer := startGrpcServer()
			httpServer := startHttpServer()
			<-done
			grpcServer.GracefulStop()
			httpServer.Shutdown(context.Background())
		}()
		time.Sleep(1 * time.Second) // Give the server a second to start
		startClient()
		close(done)
	default:
		log.Fatalf("Usage: %s <server|client|all>", os.Args[0])
	}
}

// Pour lancer le serveur gRPC et HTTP, exécutez:
// go run main.go server
// Pour lancer le client, exécutez:
// go run main.go client
// Pour lancer les deux en même temps, exécutez:
// go run main.go all
// Vous devriez voir la sortie suivante:
// Date Time server listening at
// Date Time Greeting: John Doe
//
// Vous pouvez également tester le serveur avec un client gRPC généré en utilisant le fichier protos_ext.pb.go.
//
// Cela conclut le tutoriel sur la création d'un serveur gRPC simple en Go. Vous pouvez maintenant créer vos propres services gRPC pour communiquer entre les applications en utilisant RPC.
