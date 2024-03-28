package a1b2c3

import (
	"sync"
)

func ChannelWBuffer() {
	abc := make(chan struct{}, 1)
	num := make(chan struct{}, 1)
	done := make(chan struct{})
	abc <- struct{}{}

	go func() {
		for i := 65; i <= 90; i++ {
			<-abc
			// fmt.Printf("%v", string(rune(i)))
			num <- struct{}{}
		}
	}()
	go func() {
		for i := 1; i <= 26; i++ {
			<-num
			// fmt.Printf("%v", i)
			abc <- struct{}{}
		}
		done <- struct{}{}
	}()
	// time.Sleep(1 * time.Second)
	<-done
}

func ChannelWOBuffer() {
	abc := make(chan struct{})
	num := make(chan struct{})
	done := make(chan struct{})

	go func() {
		for i := 65; i <= 90; i++ {
			// fmt.Printf("%v", string(rune(i)))
			abc <- struct{}{}
			<-num
		}
	}()
	go func() {
		for i := 1; i <= 26; i++ {
			<-abc
			// fmt.Printf("%v", i)
			num <- struct{}{}
		}
		done <- struct{}{}
	}()
	// time.Sleep(1 * time.Second)
	<-done
}

func WGLock() {
	abcMux := sync.Mutex{}
	numMux := sync.Mutex{}

	numMux.Lock()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 65; i <= 90; i++ {
			abcMux.Lock()
			// fmt.Printf("%v", string(rune(i)))
			numMux.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i++ {
			numMux.Lock()
			// fmt.Printf("%v", i)
			abcMux.Unlock()
		}
	}()
	wg.Wait()
}
