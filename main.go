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

func work(b Barrier, ch chan struct{}) {
	fmt.Print("start")
	b.Wait()
	fmt.Print("end")
	ch <- struct{}{}
}

func main() {

	b := NewBarrier(2)
	ch := make(chan struct{})
	for i := 0; i < 2; i++ {
		go work(b, ch)
	}
	for i := 0; i < 2; i++ {
		<-ch
	}
	for i := 0; i < 2; i++ {
		go work(b, ch)
	}
	for i := 0; i < 2; i++ {
		<-ch
	}
	time.Sleep(time.Second)
}
