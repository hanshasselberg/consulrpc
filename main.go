package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/consul/agent/structs"
	"github.com/i0rek/consulrpc/rpc"
)

func main() {
	args := structs.KeyRequest{
		Datacenter: "dc1",
	}
	var out structs.IndexedDirEntries
	codec, err := rpc.Codec("127.0.0.1:8300")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer codec.Close()
	err = rpc.RPC(codec, "KVS.List", &args, &out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", out)
}
