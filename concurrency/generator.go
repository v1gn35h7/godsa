package concurrency

func Generator(input []int) chan int {

	out := make(chan int)

	go func() {
		defer close(out)
		for _, v := range input {
			out <- v
		}
	}()

	return out

}
