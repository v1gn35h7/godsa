package concurrency

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicOperations() {
	var x int64
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			atomic.AddInt64(&x, 1)

		}()

	}
	wg.Wait()

	fmt.Println(x)

}
