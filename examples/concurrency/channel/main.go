package main

import "fmt"

func main() {
	var ch = make(chan int, 4)
	for i := 0; i < 10; i++ {
		select {
		case ch <- i:
		case v := <-ch:
			fmt.Println(v)
		}
	}

	// for v := range ch {
	// 	fmt.Println(v)
	// }
}
