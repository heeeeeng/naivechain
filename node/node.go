package node

import (

	"github.com/heeeeeng/naivechain/og"
	"github.com/heeeeeng/naivechain/p2p"
	"github.com/heeeeeng/naivechain/rpc"
)

type Node struct {
	og		*og.Og

	p2pServer *p2p.P2PServer
	rpcServer *rpc.RpcServer
}

func NewNode() *Node {
	n := &Node{
		p2pServer: p2p.NewP2PServer("p2p port"),
		rpcServer: rpc.NewRpcServer("rpc port"),
	}
	return n
}

func (n *Node) Start() {
	n.p2pServer.Start()
	n.rpcServer.Start()

	n.og.Start()
}

