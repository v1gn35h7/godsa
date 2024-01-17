package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

func TestTeeChannel(t testing.T) {
	i := []int{5, 7, 12, 9, 3, 8, 432, 0, 4}
	g1 := Generator(i)

	out1, out2 := Tee(g1)

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
