package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(len(nums))

	for _, n := range nums {
		go func(x int) {
			defer wg.Done()
			fmt.Printf("%d squared = %d\n", x, x*x)
		}(n)
	}

	wg.Wait()
}
