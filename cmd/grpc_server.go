package cmd

import (
	"URLShorter/internal/service"
	"URLShorter/internal/storage"
	"fmt"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedURLShortenerServer
	service service.URLService
}

func (s *server) CreateShortURL(ctx context.Context, req *proto.CreateShortURLRequest) (*proto.CreateShortURLResponse, error) {
	shortCode, err := s.service.CreateShortURL(req.GetOriginalUrl())
	if err != nil {
		return nil, err
	}

	return &proto.CreateShortURLResponse{
		ShortUrl: "http://localhost:8080/" + shortCode,
	}, nil
}

func (s *server) GetOriginalURL(ctx context.Context, req *proto.GetOriginalURLRequest) (*proto.GetOriginalURLResponse, error) {
	originalURL, err := s.service.GetOriginalURL(req.GetShortCode())
	if err != nil {
		return nil, err
	}

	return &proto.GetOriginalURLResponse{
		OriginalUrl: originalURL,
	}, nil
}

func main() {
	store, err := storage.NewRedisStore("localhost:6379")
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	urlService := service.NewURLService(store)
	grpcServer := grpc.NewServer()
	proto.RegisterURLShortenerServer(grpcServer, &server{service: urlService})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("gRPC server started on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
