package main

import (
	"fmt"
	"sync"
)


var wg = sync.WaitGroup{}

func main() {
	//ch := make(chan int) // a channel with a static type - this is an unbuffered channel, add a second parameter to set the buffer size
	ch := make(chan int,10) // buffered channel with size 10

	// Example 1 - unbuffered use case
	wg.Add(2)
	go func ()  { // this function gets an integer from the channel then prints it
		i:= <- ch
		fmt.Println(i)
		wg.Done()
	}()
	go func ()  { // this function sends an integer to the channel
		i:= 42
		ch <- i
		i = 27
		wg.Done()
	}()
	wg.Wait()
	
	// Example 2 - buffered use case 1
	wg.Add(2)
	go func (ch <-chan int)  { // exclusive receiver function from channel
		for i:= range ch {
			fmt.Println(i) // getting all the elements in the channel
		}
		wg.Done()
	}(ch)
	go func (ch chan<- int)  { // exclusive sender function to channel
		ch <- 43
		ch <- 21 // this is an unbuffered channel, need to change it to a buffered channel to run
		close(ch) // closing the channel to limit its size before someone else uses it properly
		wg.Done()
	}(ch)
	wg.Wait()
	
	// Example 3 - buffered use case 2 with a better receiver function
	ch2 := make(chan int, 15)
	wg.Add(2)
	go func (ch <-chan int)  { // exclusive receiver function from channel
		for {
			if i, ok := <- ch; ok {
				fmt.Println(i) 
			} else{
				break
			}
		}
		wg.Done()
	}(ch2)
	go func (ch chan<- int)  { // exclusive sender function to channel
		ch <- 45
		ch <- 28 // this is an unbuffered channel, need to change it to a buffered channel to run
		close(ch) // closing the channel to limit its size before someone else uses it properly
		wg.Done()
	}(ch2)
	wg.Wait()
}