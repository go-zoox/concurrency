package concurrency

import (
	"log"
	"sync"
	"time"
)

// Concurrency is a concurrent object
type Concurrency struct {
	ch chan struct{}
	wg *sync.WaitGroup
}

// New creates a new concurrency object
func New(limit int) *Concurrency {
	return &Concurrency{
		ch: make(chan struct{}, limit),
		wg: new(sync.WaitGroup),
	}
}

// Add adds a new task to the concurrency object
func (c *Concurrency) Add(task func(args ...interface{}), args ...interface{}) {
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

// Done removes a task from the concurrency object
func (c *Concurrency) Done() {
	<-c.ch
	c.wg.Done()
}

// Wait waits for all tasks to be done
func (c *Concurrency) Wait() {
	c.wg.Wait()
}
