package main

import (
	"fmt"

	"github.com/zchelalo/sa_api_gateway/internal/server"
	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/util"
)

func main() {
	bootstrap.InitLogger()

	logger := bootstrap.GetLogger()

	config, err := util.LoadConfig(".")
	if err != nil {
		logger.Fatal("cannot load config:", err)
	}

	address := fmt.Sprintf("0.0.0.0:%d", config.Port)

	bootstrap.InitRESTClient()

	services := map[constants.GRPCConstants]string{
		constants.UserMicroserviceDomain:            config.UserMicroserviceDomain,
		constants.AuthMicroserviceDomain:            config.AuthMicroserviceDomain,
		constants.ClassManagementMicroserviceDomain: config.ClassManagementMicroserviceDomain,
		constants.MemberMicroserviceDomain:          config.ClassManagementMicroserviceDomain,
	}

	for name, addr := range services {
		err := bootstrap.InitGRPCClient(addr, name)
		if err != nil {
			logger.Fatalf("cannot init grpc client for %s: %v", name, err)
		}
	}

	s := server.NewServer(address)
	s.Start()
}
