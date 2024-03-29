package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	id       int
	username string
}

func (u User) Get() int {
	return u.id
}

type I interface {
	Get() int
}

var pool = sync.Pool{
	New: func() interface{} {
		return User{1, "vicky"}
	},
}

func main1() {
	// for i := 0; i < 1000; i++ {
	// 	usr := pool.Get()

	// 	fmt.Println(usr.(I))
	// }

	p := &User{}
	fmt.Printf("%p", p)
	fmt.Println("\n")

	p1 := new(User)
	fmt.Println(p1)

	ch := make(chan int)

	go func() {
		//ch <- 1
		//ch <- 2
	}()

	<-ch

}

func main2() {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}

	//get first result
	fmt.Println("result:", <-ch)
	//do other work
	time.Sleep(2 * time.Second)
}

func main_2(2) {
	fmt.Println(waysToSplitArray([]int{1, 1, 1, 2, 3}))
}

func waysToSplitArray(nums []int) int {
	totalSum := 0
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}
	currentSum := 0

	validArrCounter := 0
	for i := 0; i < len(nums)-1; i++ {
		currentSum += nums[i]
		rsum := totalSum - currentSum
		fmt.Println(i, currentSum, rsum, totalSum)

		if currentSum >= rsum {
			validArrCounter++
		}
	}
	return validArrCounter
}
