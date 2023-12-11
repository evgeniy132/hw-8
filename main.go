package main

import (
	"fmt"
	"sync"
	"time"
)

func NewBarrier(N int) Barrier {
	var _wg sync.WaitGroup
	_wg.Add(N)
	return Barrier{
		wg: &_wg,
	}

}

type Barrier struct {
	wg *sync.WaitGroup
}

func (wgb *Barrier) Wait() {
	wgb.wg.Done()
	wgb.wg.Wait()
}

func work(b Barrier) {
	fmt.Print("start")
	b.Wait()
	fmt.Print("end")
}

func main() {
	b := NewBarrier(2)

	go work(b)
	go work(b)
	time.Sleep(time.Second)
}
