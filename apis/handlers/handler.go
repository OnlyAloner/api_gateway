package handlers

import (
	"encoding/json"

	"github.com/OnlyAloner/api_gateway/config"
	client "github.com/OnlyAloner/api_gateway/pkg/grpc_client"
	"github.com/OnlyAloner/api_gateway/pkg/logger"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
)

type handlerV1 struct {
	log        logger.Logger
	grpcClient *client.Grpcservices
	cfg        config.Config
	redis      redis.Conn
}

type HandlerV1Config struct {
	Logger     logger.Logger
	GrpcClient *client.Grpcservices
	Cfg        config.Config
	Redis      redis.Conn
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:        c.Logger,
		grpcClient: c.GrpcClient,
		cfg:        c.Cfg,
		redis:      c.Redis,
	}
}

func ProtoToStruct(s interface{}, p proto.Message) error {
	var jm jsonpb.Marshaler

	jm.EmitDefaults = true
	jm.OrigName = true

	ms, err := jm.MarshalToString(p)

	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(ms), &s)

	return err
}
