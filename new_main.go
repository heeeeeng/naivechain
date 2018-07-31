package main

import (
	"github.com/heeeeeng/naivechain/core"
	"github.com/heeeeeng/naivechain/p2p"
	"github.com/heeeeeng/naivechain/rpc"
)

func newMain() {
	bc := &core.BlockChain{}

	p2pServer := p2p.NewP2PServer(bc)
	rpcServer := rpc.NewRpcServer(bc)

	bc.Init(p2pServer, rpcServer)
	bc.Start()

}
