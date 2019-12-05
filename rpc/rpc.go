package rpc

import (
	"net"
	"net/rpc"

	msgpackrpc "github.com/hashicorp/net-rpc-msgpackrpc"
)

func Codec(dst string) (rpc.ClientCodec, error) {
	conn, err := net.Dial("tcp", dst)
	if err != nil {
		return nil, err
	}

	// 0 is the bit for consulrpc, there are others which are not needed
	// atm and they need to stay the same for backwards compat.
	// https://github.com/hashicorp/consul/blob/681767622322ae9a31b35e93bfb8d4c54a621509/agent/pool/conn.go#L7-L19.
	if _, err := conn.Write([]byte{byte(0)}); err != nil {
		conn.Close()
		return nil, err
	}
	return msgpackrpc.NewClientCodec(conn), nil
}

func RPC(codec rpc.ClientCodec, method string, args interface{}, reply interface{}) error {
	return msgpackrpc.CallWithCodec(codec, method, args, reply)
}
