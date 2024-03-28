package fizzbuzzmultithreaded

import (
	"fmt"
	"sync"
)

var (
	wg = sync.WaitGroup{}
)

// 時間複雜 O(), 空間複雜 O()
func FizzBuzz(n int) {
	fb := NewFizzBuzz()

	wg.Add(4)

	go fb.fizz()
	go fb.buzz()
	go fb.fizzbuzz()
	go fb.number()

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fb.fizzbuzzCh <- struct{}{}
		} else if i%3 == 0 {
			fb.fizzCh <- struct{}{}
		} else if i%5 == 0 {
			fb.buzzCh <- struct{}{}
		} else {
			fb.numberCh <- i
		}
		<-fb.orderCh
	}
	fb.done <- struct{}{}
	fb.done <- struct{}{}
	fb.done <- struct{}{}
	fb.done <- struct{}{}
	wg.Wait()
}

type FizzBuzzStruct struct {
	numberCh   chan int
	fizzCh     chan struct{}
	buzzCh     chan struct{}
	fizzbuzzCh chan struct{}
	orderCh    chan struct{}
	done       chan struct{}
}

func NewFizzBuzz() *FizzBuzzStruct {
	return &FizzBuzzStruct{
		numberCh:   make(chan int),
		fizzCh:     make(chan struct{}),
		buzzCh:     make(chan struct{}),
		fizzbuzzCh: make(chan struct{}),
		orderCh:    make(chan struct{}),
		done:       make(chan struct{}, 4),
	}
}

func (fb *FizzBuzzStruct) fizz() {
	defer wg.Done()
	for {
		select {
		case <-fb.fizzCh:
			fmt.Print("fizz")
			fb.orderCh <- struct{}{}
		case <-fb.done:
			return
		}
	}
}

func (fb *FizzBuzzStruct) buzz() {
	defer wg.Done()
	for {
		select {
		case <-fb.buzzCh:
			fmt.Print("buzz")
			fb.orderCh <- struct{}{}
		case <-fb.done:
			return
		}
	}
}

func (fb *FizzBuzzStruct) fizzbuzz() {
	defer wg.Done()
	for {
		select {
		case <-fb.fizzbuzzCh:
			fmt.Print("fizzbuzz")
			fb.orderCh <- struct{}{}
		case <-fb.done:
			return
		}
	}
}

func (fb *FizzBuzzStruct) number() {
	defer wg.Done()
	for {
		select {
		case v := <-fb.numberCh:
			fmt.Print(v)
			fb.orderCh <- struct{}{}
		case <-fb.done:
			return
		}
	}
}
