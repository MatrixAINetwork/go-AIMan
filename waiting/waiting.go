package waiting

import (
	"github.com/MatrixAINetwork/go-AIMan/dto"
	"github.com/MatrixAINetwork/go-AIMan/manager"
	"time"
)

type WaitInterface interface {
	Done() chan struct{}
	Quit()
	Waiting()
}
type WaitTime struct {
	wt   *time.Timer
	done chan struct{}
}

func NewWaitTime(wt time.Duration) *WaitTime {
	return &WaitTime{
		time.NewTimer(wt),
		make(chan struct{}),
	}
}
func (w *WaitTime) Done() chan struct{} {
	return w.done
}
func (w *WaitTime) Waiting() {
	go func() {
		for {
			select {
			case <-w.wt.C:
				w.done <- struct{}{}
				return
			}
		}
	}()
}
func (w *WaitTime) Quit() {
	w.wt.Stop()
}

type Web3Waiting struct {
	done   chan struct{}
	quit   chan struct{}
	waitCh chan struct{}
	Web3   *manager.Manager
}

func (w *Web3Waiting) makeChan() {
	w.done = make(chan struct{}, 1)
	w.quit = make(chan struct{}, 1)
	w.waitCh = make(chan struct{}, 1)
}
func (w *Web3Waiting) Done() chan struct{} {
	return w.done
}
func (w *Web3Waiting) Quit() {
	select {
	case w.quit <- struct{}{}:
	default:
	}
}
func (w *Web3Waiting) WaitingFn(fn func()) {
	go func() {
		for {
			select {
			case <-time.After(100 * time.Millisecond):
				go func() {
					w.waitCh <- struct{}{}
					fn()
					<-w.waitCh
				}()
			case <-w.quit:
				return
			}
		}
	}()
}

type WaitBlockHeight struct {
	Web3Waiting
	height uint64
	Result uint64
}

func NewWaitBlockHeight(web *manager.Manager, height uint64) *WaitBlockHeight {
	w := &WaitBlockHeight{
		Web3Waiting{Web3: web},
		height, 0,
	}
	w.makeChan()
	return w
}
func (w *WaitBlockHeight) Waiting() {
	w.WaitingFn(func() {
		bm, err := w.Web3.Man.GetBlockNumber()
		if err == nil {
			if bm.Uint64() >= w.height {
				w.Result = bm.Uint64()
				w.done <- struct{}{}
				w.Quit()
			}
		}

	})
}

type WaitTxReceipt struct {
	Web3Waiting
	txhash  string
	Receipt *dto.TransactionReceipt
}

func NewWaitTxReceipt(web *manager.Manager, txHash string) *WaitTxReceipt {
	w := &WaitTxReceipt{
		Web3Waiting{Web3: web},
		txHash, nil,
	}
	w.makeChan()
	return w
}
func (w *WaitTxReceipt) Waiting() {
	w.WaitingFn(func() {
		receipt, err := w.Web3.Man.GetTransactionReceipt(w.txhash)
		if err == nil {
			if receipt != nil && receipt.BlockNumber != nil && receipt.BlockNumber.Uint64() > 0 {
				w.Receipt = receipt
				w.done <- struct{}{}
				w.Quit()
			}
		}
	})
}

type WaitTxReceiptAry struct {
	Web3Waiting
	Receipts map[string]*dto.TransactionReceipt
	waiting  int
}

func NewWaitTxReceiptAry(web *manager.Manager, txHashAry []string) *WaitTxReceiptAry {
	wt := &WaitTxReceiptAry{
		Web3Waiting{Web3: web},
		make(map[string]*dto.TransactionReceipt, len(txHashAry)),
		len(txHashAry),
	}
	wt.makeChan()
	for _, item := range txHashAry {
		wt.Receipts[item] = nil
	}
	return wt
}
func (wt *WaitTxReceiptAry) Waiting() {
	wt.WaitingFn(func() {
		errNum := 0
		for key, item := range wt.Receipts {
			if item != nil {
				continue
			}
			receipt, err := wt.Web3.Man.GetTransactionReceipt(key)
			if err == nil && receipt != nil && receipt.BlockNumber != nil && receipt.BlockNumber.Uint64() > 0 {
				wt.Receipts[key] = receipt
				wt.waiting--
				if wt.waiting <= 0 {
					wt.done <- struct{}{}
					wt.Quit()
				}
			} else {
				errNum++
				if errNum > 5 {
					return
				}
			}
		}
	})
}

type WaitMinerNumChange struct {
	Web3Waiting
	Number uint64
	Result uint64
}

func NewWaitMinerNumChange(web *manager.Manager, number uint64) *WaitMinerNumChange {

	w := &WaitMinerNumChange{
		Web3Waiting{Web3: web},
		number, 0,
	}
	w.makeChan()
	return w
}
func (w *WaitMinerNumChange) Waiting() {
	w.WaitingFn(func() {
		bm, _ := w.Web3.Man.GetBlockNumber()
		tmp, err := w.Web3.Man.GetTopologyStatusByNumber(bm)
		if err == nil {
			if uint64(len(tmp.Miners)) == w.Number {
				w.Result = uint64(len(tmp.Miners))
				w.done <- struct{}{}
				w.Quit()
			}
		}
	})
}

type WaitVManStrDisappear struct {
	Web3Waiting
	Manstr string
	IsHere bool
}

func NewWaitVManStrDisappear(web *manager.Manager, manstr string) *WaitVManStrDisappear {

	w := &WaitVManStrDisappear{
		Web3Waiting{Web3: web},
		manstr, true,
	}
	w.makeChan()
	return w
}

func (w *WaitVManStrDisappear) Waiting() {
	w.WaitingFn(func() {
		bm, _ := w.Web3.Man.GetBlockNumber()
		tmp, err := w.Web3.Man.GetTopologyStatusByNumber(bm)
		if err == nil {

			for _, value := range tmp.Validators {
				if value.Account == w.Manstr {
					w.IsHere = true
					break
				} else {
					w.IsHere = false
				}
			}
			if !w.IsHere {
				w.done <- struct{}{}
				w.Quit()
			}
		}
	})
}
