package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.Lock()
	c.count++
	c.Unlock()
}

func (c *Counter) Count() uint64 {
	c.Lock()
	defer c.Unlock()
	return c.count
}

func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}

func main() {
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go worker(&counter, &wg)
	}

	wg.Wait()
	fmt.Println(counter.Count())
}
