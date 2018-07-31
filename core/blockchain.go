package core

import (
	"github.com/heeeeeng/naivechain/core/types"
)

type BlockChain struct {
	genesisBlock *types.Block
	latestBlock  *types.Block

	chain	[]*types.Block

	p2pServer Server
	rpcServer Server
}

type Server interface {
	Start()
}

func (bc *BlockChain) Init(p2pServer, rpcServer Server) {
	bc.p2pServer = p2pServer
	bc.rpcServer = rpcServer
}

func (bc *BlockChain) Start() {
	bc.p2pServer.Start()
	bc.rpcServer.Start()
}

func (bc *BlockChain) GetBlock(i int) *types.Block {
	if i == -1 {
		return bc.chain[len(bc.chain)-1]
	}
	if i < 0 || i > len(bc.chain) {
		return nil
	}
	return bc.chain[i]
}

func (bc *BlockChain) String() string {
	result := ""
	for _, block := range bc.chain{
		result += block.String()
	}
	return result
}
func (bc *BlockChain) Len() int { return int(bc.latestBlock.Index) }
