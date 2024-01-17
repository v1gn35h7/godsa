package main

import (
	"fmt"
	"sync"

	"github.com/vicky1115050/godsa/concurrency"
)

func main() {
	g1 := concurrency.Generator([]int{5, 7, 12, 9, 3, 8, 432, 0, 4})

	out1, out2 := concurrency.Tee(g1)

	var wg sync.WaitGroup

	wg.Add(2)

	go process(out1, 1, &wg)
	go process(out2, 2, &wg)

	wg.Wait()

}

func process(in chan int, i int, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range in {
		fmt.Println("Recived from -> ", i, "Val: ", v)
	}

}
