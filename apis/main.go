package api

import (
	"net/http"

	_ "github.com/OnlyAloner/api_gateway/apis/docs" // for api
	"github.com/OnlyAloner/api_gateway/apis/handlers"
	"github.com/OnlyAloner/api_gateway/config"
	client "github.com/OnlyAloner/api_gateway/pkg/grpc_client"
	"github.com/OnlyAloner/api_gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	//"github.com/OnlyAloner/api_gateway/apis/docs"
)

// Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *client.Grpcservices
	Cfg        config.Config
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	handler := handlers.New(&handlers.HandlerV1Config{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "api OK"})
	})

	r.POST("handlers/post_service/create", handler.Create)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
