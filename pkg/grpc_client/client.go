package client

import (
	"fmt"

	"github.com/OnlyAloner/api_gateway/config"
	"github.com/OnlyAloner/api_gateway/proto/post_service"

	"google.golang.org/grpc"
)

type GrpcservicesI interface {
	PostService() post_service.PostServiceClient
}

type Grpcservices struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*Grpcservices, error) {
	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceNewHost, cfg.PostServiceNewPort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("sms service dial host: %s port:%d err: %s",
			cfg.PostServiceNewHost, cfg.PostServiceNewPort, err)
	}
	return &Grpcservices{
		cfg: cfg,
		connections: map[string]interface{}{
			"post_service": post_service.NewPostServiceClient(connPost),
		},
	}, nil
}

func (g *Grpcservices) PostService() post_service.PostServiceClient {
	return g.connections["post_service"].(post_service.PostServiceClient)
}
