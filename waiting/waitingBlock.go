package waiting

import (
	"math/big"
	"github.com/matrix/go-AIMan/common"
)
type BlockFilter interface {
	HandleBlock(block *common.Block)
}
type WaitingBlock struct {
	Web3Waiting
	startheight uint64
	endheight   uint64
	Filter []BlockFilter
}
func(w* WaitingBlock)Waiting(){
	w.WaitingFn(func() {
		bm,err := w.Web3.Man.GetBlockByNumber(new(big.Int).SetUint64(w.startheight),true)
		if err == nil{
			for _,item := range w.Filter {
				item.HandleBlock(bm)
			}
			w.startheight++
			if w.startheight>w.endheight{
				w.done <- struct{}{}
				w.Quit()
			}
		}

	})
}
type TransactionFilter interface{
	HandleTx(transaction *common.RPCTransaction)
}
type TxHandle struct {
	Filter []TransactionFilter
}
func (Tf* TxHandle)HandleBlock(block *common.Block){
	for i:=0;i< len(block.Transactions);i++ {
		for _,item := range Tf.Filter{
			item.HandleTx(block.Transactions[i])
		}
	}
}