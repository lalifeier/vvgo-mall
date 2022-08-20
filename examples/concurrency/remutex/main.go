package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.RWMutex
	count uint64
}

func (c *Counter) Incr() {
	c.Lock()
	c.count++
	c.Unlock()
}

func (c *Counter) Count() uint64 {
	c.RLock()
	defer c.RUnlock()
	return c.count
}

func main() {
	var counter Counter

	for i := 0; i < 10; i++ {
		go func() {
			for {
				fmt.Println(counter.Count())
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for {
		counter.Incr()
		time.Sleep(time.Second)
	}
}
