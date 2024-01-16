package pro

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
)

func TestMa(test *testing.T) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	// Start broker
	br := NewBroker()

	var wg sync.WaitGroup

	wg.Add(9)
	// Start 5 producer
	t := fmt.Sprintf("%s%d", "topic", 1)
	br.CreateTopic(t)
	for i := 0; i < 5; i++ {

		go func(p int) {
			defer wg.Done()
			fmt.Println("Producer: ", p, " started ...")

			for j := 0; j < 10; j++ {
				br.Publish(t, j)
				//	fmt.Println("Producer: ", p, " produced -->", j, " on topic -->", t)

			}

		}(i)

	}

	// Start 5 c
	for i := 0; i < 4; i++ {

		go func(c int) {
			defer wg.Done()

			msgcount := 0

			fmt.Println("Consumer: ", c, " started ...")
			t := fmt.Sprintf("%s%d", "topic", 1)
			r, p, err := br.Subscribe(strconv.Itoa(c), t)

			if err != nil {
				fmt.Println("Consumer --> ", c, " stopped with err: ", err)
				return
			}

			if r == nil {
				fmt.Println("Consumer --> ", c, " stopped with err: ", "Nil channel")
				return
			}

			// Wait until broker allows to consume
			<-r

			fmt.Println("Consumer --> ", c, " assigned with partition: ", p)

			for {

				r, o, err := br.Consume(t, strconv.Itoa(c), p)

				if err != nil {
					fmt.Println("Consumer --> ", c, " stopped with err: ", err)
					break
				}

				if r == 0 && o == 0 {
					continue
				}

				fmt.Println("Consumer --> ", c, "Consumed --> ", r, " from topic -->", 1, " at offset --> ", o)
				msgcount += 1

				if msgcount >= 10 {
					br.UnSubscribe(strconv.Itoa(c), t, p)
					break
				}

			}

		}(i)

	}

	wg.Wait()

	br.Flush()

}
