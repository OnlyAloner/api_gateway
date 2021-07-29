package api

import (
	"crudprojforapi/api_gateway/config"
	"crudprojforapi/api_gateway/pkg/logger"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Logger     logger.Logger
	GrpcClient *client.grpcservices
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

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "api OK"})
	})
	r.Run(":8080")

}
