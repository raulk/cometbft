package mempool

import (
	"sync/atomic"

	"github.com/cometbft/cometbft/types"
)

// mempoolTx is an entry in the mempool
type MempoolTx struct { //nolint: revive // The variable name stutters; this is only an ad-hoc hack to access this struct from the test/e2e/fast-prototyping.
	height    int64    // height that this tx had been validated in
	gasWanted int64    // amount of gas this tx states it will require
	tx        types.Tx // validated by the application
}

// Height returns the height for this transaction
func (memTx *MempoolTx) Height() int64 {
	return atomic.LoadInt64(&memTx.height)
}

func (memTx *MempoolTx) GetTx() types.Tx {
	return memTx.tx
}
