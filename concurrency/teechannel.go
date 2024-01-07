package concurrency

import "fmt"

func Tee(in chan int) (chan int, chan int) {
	out1 := make(chan int)
	out2 := make(chan int)

	go func() {
		for v := range in {
			for i := 0; i < 2; i++ {
				select {
				case out1 <- v:
					fmt.Println("Sent to out1")
				case out2 <- v:
					fmt.Println("Sent to out2")
				}
			}
		}
		close(out1)
		close(out2)
	}()

	return out1, out2

}
