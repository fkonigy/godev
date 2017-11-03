package main

import (
	"fmt"
)

type Counter struct {
	c    chan int
	done chan struct{}
	i    int
}

func NewCounter() *Counter {
	counter := new(Counter)
	counter.c = make(chan int)
	counter.done = make(chan struct{})

	go func() {
		for {
			select {
			case counter.c <- counter.i:
				counter.i += 1
			case <-counter.done:
				return
			}
		}
	}()

	return counter
}

// make it a readonly chan
func (c *Counter) GetSource() <-chan int {
	return c.c
}

func (c *Counter) Stop() {
	c.done <- struct{}{}
}

func main() {
	c := NewCounter()
	read := c.GetSource()

	fmt.Printf("%d\n", <-read)
	// read <- 1 # it is a readonly chan
	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	// done <- struct{}{}
	c.Stop()

	fmt.Printf("%d\n", <-read) // it is close now
}
