package main

import (
	"fmt"

	"github.com/vicky1115050/godsa/algo"
)

func main() {
	fmt.Println("################### GO - DSA #########################")
	list := make([]algo.KnapsackItem, 0)
	list = append(list, algo.KnapsackItem{P: 60, W: 10})
	list = append(list, algo.KnapsackItem{P: 100, W: 20})
	list = append(list, algo.KnapsackItem{P: 120, W: 30})
	r := algo.FractionalKanpsack(50, list)
	fmt.Println(r)
}
