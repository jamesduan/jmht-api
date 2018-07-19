package rpc

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

func JsonRpcClient(network, address string, timeout time.Duration) (*rpc.Client, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), err
}
