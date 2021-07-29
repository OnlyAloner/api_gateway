package main

import (
	api "crudprojforapi/api_gateway/apis"
	"crudprojforapi/api_gateway/pkg/logger"
	"sync"

	config "crudprojforapi/api_gateway"
	"crudprojforapi/api_gateway/pkg/grpc_client/client"

	_ "github.com/lib/pq"
)

var (
	log        logger.Logger
	cfg        config.Config
	grpcClient *client.GrpcStaff
	err        error
)

func main() {

	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "api_gateway")

	grpcClient, err = client.New(cfg)
	if err != nil {
		log.Error("grpc dial error", logger.Error(err))
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		server := api.New(api.Config{
			Logger:     log,
			GrpcClient: grpcClient,
			Cfg:        cfg,
		})
		err := server.Run(cfg.HttpPort)
		if err != nil {
			panic(err)
		}
		panic("api server has finished")
	}(&wg)

	wg.Wait()
}
