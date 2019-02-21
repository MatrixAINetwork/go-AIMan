package waiting

import "reflect"

type MultiWaiting struct {
	WaitArray []WaitInterface
	waitCase  []reflect.SelectCase
}

func NewMultiWaiting(args ...WaitInterface) *MultiWaiting {
	return &MultiWaiting{
		args,
		nil,
	}
}
func (mw *MultiWaiting) Waiting() int {
	mw.waitCase = make([]reflect.SelectCase, len(mw.WaitArray))
	for i, item := range mw.WaitArray {
		item.Waiting()
		mw.waitCase[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(item.Done())}
	}
	chosen, _, _ := reflect.Select(mw.waitCase)
	for _, item := range mw.WaitArray {
		item.Quit()
	}
	return chosen
}
