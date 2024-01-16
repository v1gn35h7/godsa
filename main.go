package main

import (
	"fmt"
	"sync"
	"sync"

	"github.com/vicky1115050/godsa/algo"
)

func main() {
	fmt.Println("################### GO - DSA #########################")
	//N = 3, W = 4, profit[] = {1, 2, 3}, weight[] = {4, 5, 1}
	fmt.Println(algo.KnapSack01(3, 6, []int16{1, 2, 3}, []int16{10, 15, 40}))
}
