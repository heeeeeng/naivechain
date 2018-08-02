package rpc 

import (
	"log"
	"net/http"

	"github.com/heeeeeng/naivechain/core"
)

type RpcServer struct {
	rpcport		string

	blockchain *core.BlockChain
}

func NewRpcServer(rpcport string, bc *core.BlockChain) *RpcServer {
	srv := &RpcServer{
		rpcport:	rpcport,
		blockchain: bc,
	}
	return srv
}

func (srv *RpcServer) Start() {
	// http.HandleFunc("/blocks", handleBlocks)
	// http.HandleFunc("/add_peer", handleAddPeer)
	go func() {
		log.Println("Listen HTTP on",  srv.rpcport)
		log.Fatalln("start api server", http.ListenAndServe(srv.rpcport, nil))
	}()
}
