package algo

import (
	"fmt"
)

func DistinctElementsInWindow(list []int, k int) {
	hm := make(map[int]int)
	n := len(list)

	for i := 0; i < k; i++ {
		hm[list[i]] = hm[list[i]] + 1
	}

	fmt.Println(len(hm))

	for j := k; j < n; j++ {
		if hm[list[j-k]] == 1 {
			delete(hm, list[j-k])
		} else {
			hm[list[j-k]] = hm[list[j-k]] - 1
		}
		hm[list[j]] = 1
		fmt.Println(len(hm))
	}

}
