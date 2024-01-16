package concurrency

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func TstContext() {

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(cx context.Context) {
			defer wg.Done()
			for {
				select {
				case <-cx.Done():
					fmt.Printf("-")
					return
				default:
					fmt.Printf(".")
				}
			}

		}(ctx)

	}

	go func() {
		time.Sleep(1000)
		fmt.Println("Cancelling context")
		cancel()
	}()

	wg.Wait()
}
