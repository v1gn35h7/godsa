package main

import (
	"fmt"
	"sync"

	"github.com/vicky1115050/godsa/conc/pro"
)

func main() {
	fmt.Println("Test LRU...")

	cache := pro.NewCacheStore()

	key := "4f6s4f65s"

	var wg sync.WaitGroup

	// setters
	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			k := fmt.Sprintf("%s%d", key, x)
			cache.Set(k, x+100)
			fmt.Println("S -->", x, "set v -->", x+100)
		}(i)
	}

	// getters
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			v, err := cache.Get(fmt.Sprintf("%s%d", key, x))
			if err != nil {
				fmt.Println("Failed to read")
				return
			}
			fmt.Println("G -->", x, "get v -->", v)
		}(i)
	}

	wg.Wait()

	cache.Flush()
}
