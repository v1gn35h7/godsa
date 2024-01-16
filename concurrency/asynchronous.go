package concurrency

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"golang.org/x/sync/errgroup"
)

func AsyncApi(x int) <-chan int {
	out := make(chan int)

	go func(y int) {
		for i := 0; i < 100; i++ {
			y += 1
		}
		out <- y

	}(x)

	return out
}

func AsyncCall() {

	var a int
	a = 1

	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		time.Sleep(10000)
		a = <-AsyncApi(1)

		return nil
	})

	go func() {
		for {
			fmt.Printf("%d->", a)
		}
	}()

	fmt.Printf("\n")

	err := g.Wait()

	slog.Error("Failed", err)
}
