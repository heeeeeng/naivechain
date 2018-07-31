package rpc 

import (
	"github.com/heeeeeng/naivechain/core"
)

type RpcServer struct {
	blockchain *core.BlockChain
}

func NewRpcServer(bc *core.BlockChain) *RpcServer {
	srv := &RpcServer{
		blockchain: bc,
	}
	return srv
}

func (srv *RpcServer) Start() {

}