package algo

import "fmt"

/*
 * Merge Sort Algorith
 */
func MergeSort(list []int64) []int64 {
	if len(list) == 1 {
		return list
	}

	left := MergeSort(list[:len(list)/2])
	right := MergeSort(list[(len(list) / 2):])
	return Merge(left, right)
}

func Merge(list1 []int64, list2 []int64) []int64 {
	var res []int64
	var li, ri int
	for li < len(list1) && ri < len(list2) {
		if list1[li] < list2[ri] {
			res = append(res, list1[li])
			li++
		} else {
			res = append(res, list2[ri])
			ri++
		}
	}

	if li < len(list1) {
		res = append(res, list1[li:]...)
	} else {
		res = append(res, list2[ri:]...)
	}
	return res
}

/*
* Heap Sort Algorithim
 */
func HeapSort(input []int) []int {

	m := len(input) / 2

	for i := m - 1; i >= 0; i-- {
		fmt.Println(i)
		heapify(input, i)
	}

	for i := len(input) - 1; i > 0; i-- {
		input[i], input[0] = input[0], input[i]
		heapify(input[:i], 0)
	}

	return input
}

func heapify(input []int, i int) {

	root := i
	l := 2*i + 1
	r := 2*i + 2

	if l < len(input) && input[l] > input[root] {
		root = l
	}
	if r < len(input) && input[r] > input[root] {
		root = r
	}

	if root != i {
		input[root], input[i] = input[i], input[root]
		heapify(input, root)
	}
}
