package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/gocs/microservice_item/proto"
)

func main() {
	repo := &ItemRepository{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen:", err)
	}
	s := grpc.NewServer()

	pb.RegisterItemServiceServer(s, &service{repo})
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve:", err)
	}
}

///////////////////////////////////////////////////////////////////////////////

const (
	port = ":50051"
)

///////////////////////////////////////////////////////////////////////////////

// ItemRepository ...
type itemRepository interface {
	Create(*pb.Item) (*pb.Item, error)
}

///////////////////////////////////////////////////////////////////////////////

// ItemRepository ...
type ItemRepository struct {
	mu    sync.Mutex
	items []*pb.Item
}

// Create ...
func (ir *ItemRepository) Create(item *pb.Item) (*pb.Item, error) {
	ir.mu.Lock()
	defer ir.mu.Unlock()

	updated := append(ir.items, item)
	ir.items = updated
	return item, nil
}

///////////////////////////////////////////////////////////////////////////////

type service struct {
	repo itemRepository
}

// CreateItem ...
func (s *service) CreateItem(ctx context.Context, r *pb.Item) (*pb.Response, error) {
	item, err := s.repo.Create(r)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Id:      NewID().String(),
		Created: true,
		Item:    item,
		Message: "created item",
		Reciept: getReceipt(),
	}, nil
}

// NewID ...
func NewID() uuid.UUID {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatal("no reciept generated...")
	}
	return uuid
}

func getReceipt() string {
	return NewID().String()
}

///////////////////////////////////////////////////////////////////////////////
