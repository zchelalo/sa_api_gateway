package bootstrap

import (
	"fmt"
	"sync"

	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conns = make(map[constants.GRPCConstants]*grpc.ClientConn)
	mu    sync.RWMutex
)

func InitGRPCClient(address string, service constants.GRPCConstants) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := conns[service]; exists {
		return nil
	}

	cc, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	conns[service] = cc
	return nil
}

func GetGRPCClient(service constants.GRPCConstants) (*grpc.ClientConn, error) {
	mu.RLock()
	defer mu.RUnlock()

	conn, exists := conns[service]
	if !exists {
		return nil, fmt.Errorf("grpc client for %s not found", service)
	}

	return conn, nil
}

func CloseGRPCClients() {
	mu.Lock()
	defer mu.Unlock()

	for key, conn := range conns {
		if conn != nil {
			conn.Close()
		}
		delete(conns, key)
	}
}
