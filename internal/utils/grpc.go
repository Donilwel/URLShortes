package utils

import (
	"URLShortes/internal/service"
	"URLShortes/proto"
	"context"
)

type GRPCServer struct {
	proto.UnimplementedURLShortenerServer
	service service.URLShortener
}

func NewGRPCServer(s service.URLShortener) *GRPCServer {
	return &GRPCServer{service: s}
}

func (s *GRPCServer) Shorten(ctx context.Context, req *proto.ShortenRequest) (*proto.ShortenResponse, error) {
	shortURL, err := s.service.ShortenURL(ctx, req.Url)
	if err != nil {
		return nil, err
	}
	return &proto.ShortenResponse{ShortUrl: shortURL}, nil
}

func (s *GRPCServer) Expand(ctx context.Context, req *proto.ExpandRequest) (*proto.ExpandResponse, error) {
	originalURL, err := s.service.ExpandURL(ctx, req.ShortUrl)
	if err != nil {
		return nil, err
	}
	return &proto.ExpandResponse{OriginalUrl: originalURL}, nil
}
