package api

import (
	"net/http"

	"github.com/OnlyAloner/api_gateway/apis/handlers"
	"github.com/OnlyAloner/api_gateway/config"
	client "github.com/OnlyAloner/api_gateway/pkg/grpc_client"
	"github.com/OnlyAloner/api_gateway/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Config struct {
	Logger     logger.Logger
	GrpcClient *client.Grpcservices
	Cfg        config.Config
}

func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Static("/images", "./static/images")

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	config.AllowHeaders = append(config.AllowHeaders, "*")
	config.AllowMethods = append(config.AllowMethods, "OPTIONS")

	r.Use(cors.New(config))
	handler := handlers.New(&handlers.HandlerV1Config{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "api OK"})
	})
	r.POST("handlers/post_service/create", handler.Create)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
