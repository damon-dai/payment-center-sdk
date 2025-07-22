package payment_center_sdk

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"

	"github.com/shimingyah/pool"
	"google.golang.org/grpc/metadata"
)

var poolClient map[string]pool.Pool

func GetConnect(addr string) (pool.Conn, error) {
	p, ok := poolClient[addr]
	if !ok {
		var err error
		p, err = pool.New(addr, pool.DefaultOptions)
		if err != nil {
			log.Fatalf("failed to new pool: %v", err)
			return nil, err
		}
		if poolClient == nil {
			poolClient = make(map[string]pool.Pool)
		}
		poolClient[addr] = p
	}
	conn, err := p.Get()
	if err != nil {
		log.Fatalf("failed to get conn: %v", err)
	}
	return conn, err
}

func NewGrpcClient(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                30 * time.Second, // 心跳间隔
			Timeout:             10 * time.Second, // 心跳超时
			PermitWithoutStream: true,             // 即使没有活跃流也发送心跳
		}),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetMetadataCtx(traceId, source, token string, kv ...string) context.Context {
	wkv := []string{"trace_id", traceId, "project_source", source, "token", token}
	if len(kv) > 0 {
		wkv = append(wkv, kv...)
	}
	md := metadata.Pairs(wkv...)
	mdCtx := metadata.NewOutgoingContext(context.Background(), md)
	return mdCtx
}
