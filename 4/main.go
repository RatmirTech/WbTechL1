package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	var wg sync.WaitGroup
	numWorkers := 4
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	<-sigCh
	cancel()
	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup, _ int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
