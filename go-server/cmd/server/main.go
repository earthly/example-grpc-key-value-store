package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/earthly/example-grpc-key-value-store/go-server/kvapi"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var errKeyNotFound = fmt.Errorf("key not found")

// server is used to implement kvapi.KeyValueServer
type server struct {
	kvapi.UnimplementedKeyValueServer
	data map[string]string
}

// Set stores a given value under a given key
func (s *server) Set(ctx context.Context, in *kvapi.SetRequest) (*kvapi.SetReply, error) {
	key := in.GetKey()
	value := in.GetValue()
	log.Printf("serving set request for key %q and value %q", key, value)

	s.data[key] = value

	reply := &kvapi.SetReply{}
	return reply, nil
}

// Get returns a value associated with a key to the client
func (s *server) Get(ctx context.Context, in *kvapi.GetRequest) (*kvapi.GetReply, error) {
	key := in.GetKey()
	log.Printf("serving get request for key %q", key)

	value, ok := s.data[key]
	if !ok {
		return nil, errKeyNotFound
	}

	reply := &kvapi.GetReply{
		Value: value,
	}
	return reply, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", port)
	serverInstance := server{
		data: make(map[string]string),
	}
	s := grpc.NewServer()
	kvapi.RegisterKeyValueServer(s, &serverInstance)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
