package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)	// changes the threads to run and can change the performance
	for i:=0;i<10;i++ {
		wg.Add(2)
		m.RLock()
		go sayHello("Hello")
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello(msg string) {
	fmt.Printf("%v from %v\n",msg,counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}