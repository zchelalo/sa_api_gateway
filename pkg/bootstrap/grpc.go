package bootstrap

import (
	"sync"

	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conns = make(map[constants.GRPCConstants]*grpc.ClientConn)
	mu    sync.Mutex
)

func InitGRPCClient(address string, service constants.GRPCConstants) error {
	mu.Lock()
	defer mu.Unlock()

	if conns[service] != nil {
		return nil
	}

	cc, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	conns[service] = cc
	return nil
}

func GetGRPCClient(service constants.GRPCConstants) *grpc.ClientConn {
	mu.Lock()
	defer mu.Unlock()

	return conns[service]
}
