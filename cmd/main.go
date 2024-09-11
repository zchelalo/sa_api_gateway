package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zchelalo/sa_api_gateway/internal/server"
	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
	"github.com/zchelalo/sa_api_gateway/pkg/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	ctx := context.Background()

	logger := bootstrap.InitLogger()
	address := fmt.Sprintf("0.0.0.0:%d", config.Port)

	cc, err := bootstrap.InitGRPCClient(config.UserMicroserviceDomain)
	if err != nil {
		log.Fatal("cannot init grpc client:", err)
	}

	s := server.NewServer(ctx, logger, address, cc)
	s.Start()
}
