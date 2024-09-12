package main

import (
	"fmt"

	"github.com/zchelalo/sa_api_gateway/internal/server"
	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/util"
	"google.golang.org/grpc"
)

func main() {
	logger := bootstrap.InitLogger()
	ctx := bootstrap.InitContext()

	config, err := util.LoadConfig(".")
	if err != nil {
		logger.Fatal("cannot load config:", err)
	}

	address := fmt.Sprintf("0.0.0.0:%d", config.Port)

	conns := make(map[string]*grpc.ClientConn)

	services := map[string]string{
		string(constants.UserMicroserviceDomain): config.UserMicroserviceDomain,
	}

	for name, addr := range services {
		conn, err := bootstrap.InitGRPCClient(addr)
		if err != nil {
			logger.Fatalf("cannot init grpc client for %s: %v", name, err)
		}
		conns[name] = conn
		defer conn.Close()
	}

	s := server.NewServer(ctx, logger, address, conns)
	s.Start()
}
