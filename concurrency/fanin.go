package concurrency

import "sync"

func FanIN(inbound ...chan int) chan int {

	out := make(chan int)
	go func() {
		var wg sync.WaitGroup
		for _, c := range inbound {
			wg.Add(1)
			go func(inp chan int) {
				defer wg.Done()
				for v := range inp {
					out <- v
				}

			}(c)
		}

		wg.Wait()
		close(out)
	}()

	return out

}
