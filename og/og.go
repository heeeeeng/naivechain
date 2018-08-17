package og

import (
	"github.com/heeeeeng/naivechain/core"
	"github.com/heeeeeng/naivechain/account"
	"github.com/heeeeeng/naivechain/worker"
)

type Og struct {
	dag 	*core.Dag
	txpool	*core.TxPool

	worker			*worker.Worker
	accountManager	*account.AccountManager
	
	manager		*Manager
}

func (og *Og) Start() {}
