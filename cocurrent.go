package cocurrent

import (
	"log"
	"sync"
	"time"
)

// Cocurrent is a concurrent object
type Cocurrent struct {
	ch chan struct{}
	wg *sync.WaitGroup
}

// New creates a new cocurrent object
func New(limit int) *Cocurrent {
	return &Cocurrent{
		ch: make(chan struct{}, limit),
		wg: new(sync.WaitGroup),
	}
}

// Add adds a new task to the cocurrent object
func (c *Cocurrent) Add(task func(args ...interface{}), args ...interface{}) {
	c.wg.Add(1)
	c.ch <- struct{}{}

	go func() {
		defer c.Done()
		defer func() {
			if err := recover(); err != nil {
				// @TODO sleep before retry
				time.Sleep(time.Second)

				log.Println("retrying task:", err)
				c.Add(task, args...)
			}
		}()

		task(args...)
	}()
}

// Done removes a task from the cocurrent object
func (c *Cocurrent) Done() {
	<-c.ch
	c.wg.Done()
}

// Wait waits for all tasks to be done
func (c *Cocurrent) Wait() {
	c.wg.Wait()
}
