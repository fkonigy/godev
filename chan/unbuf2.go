package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Context struct {
	done chan struct{}
}

func NewContext() *Context {
	ctx := new(Context)
	ctx.done = make(chan struct{})
	return ctx
}

func (c *Context) Stop() {
	close(c.done)
}

func (c *Context) GetDone() <-chan struct{} {
	return c.done
}

type Counter struct {
	ctx *Context
	c   chan int
	i   int
}

func NewCounter(ctx *Counter) *Counter {
	counter := new(Counter)
	counter.c = make(chan int)
	counter.ctx = ctx

	wg.Add(1)
	go func() {
		defer wg.Done()
		done := counter.ctx.GetDone()
		for {
			select {
			case counter.c <- counter.i:
				counter.i += 1
			case <-done:
				fmt.Printf("Conter terminated.\n")
			}
		}
	}()

	return counter
}

// make it a readonly chan
func (c *Counter) GetSource() <-chan int {
	return c.c
}

func main() {
	ctx := NewContext()
	c := NewCounter(ctx)

	read := c.GetSource()

	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)
	ctx.Stop()
	wg.Wait()

}
