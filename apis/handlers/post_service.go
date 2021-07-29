package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/OnlyAloner/api_gateway/apis/models"
	"github.com/OnlyAloner/api_gateway/config"
	client "github.com/OnlyAloner/api_gateway/pkg/grpc_client"
	"github.com/OnlyAloner/api_gateway/pkg/logger"
	"github.com/OnlyAloner/api_gateway/proto/post_service"
	"github.com/gin-gonic/gin"
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

func (h *handlerV1) Create(c *gin.Context) {
	var (
		payload  post_service.CreateRequest
		response models.CreateResponse
	)
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Code:    http.StatusBadRequest,
			Message: "Error while binding json to struct",
			Reason:  err.Error(),
		})
		return
	}

	res, err := h.grpcClient.PostService().Create(
		context.Background(),
		&payload,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error while creating post",
			Reason:  err.Error(),
		})
		return
	}

	err = ProtoToStruct(&response, res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Error{
			Code:    http.StatusInternalServerError,
			Message: "Error while converting proto to struct",
			Reason:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
