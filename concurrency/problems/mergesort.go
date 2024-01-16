package problems

import (
	"github.com/vicky1115050/godsa/algo"
)

func MergeSort(input []int64) chan []int64 {
	res := make(chan []int64)
	if len(input) == 1 {
		go func() {
			res <- input
		}()
		return res
	}

	go func() {

		m := len(input) / 2

		left := MergeSort(input[:m])
		right := MergeSort(input[m:])

		res <- algo.Merge(<-left, <-right)
		close(res)
	}()

	return res
}
