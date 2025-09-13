package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(id int, ch <-chan int) {
	for v := range ch {
		fmt.Printf("worker %d: %d\n", id, v)
	}
}

func main() {
	n := 5
	if len(os.Args) > 1 {
		if v, err := strconv.Atoi(os.Args[1]); err == nil && v > 0 {
			n = v
		}
	}

	ch := make(chan int)
	for i := 0; i < n; i++ {
		go worker(i+1, ch)
	}

	i := 0
	for {
		ch <- i
		i++
		time.Sleep(100 * time.Millisecond)
	}
}
