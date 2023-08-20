package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestConcurrency(t *testing.T) {
	c := New(2)

	for i := 0; i < 4; i++ {
		fmt.Println("Adding task", i)
		c.Add(func(args ...interface{}) {
			index := args[0].(int)

			fmt.Println("task", index, time.Now())
			time.Sleep(3 * time.Second)

			// if index == 0 {
			// 	panic("panic error for task 0")
			// }

			// panic(fmt.Sprintf("%d", index))
			// c.Done()
		}, i)
	}

	c.Wait()
}
