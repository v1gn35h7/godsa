package ds

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Hash struct{}

func (hs Hash) getHashMap() map[int]string {
	hash_table := make(map[int]string)

	//hash_table[0] = "test"
	//hash_table[1] = "sonu"

	//i, ok := hash_table[1]
	//delete(hash_table, 1)

	//fmt.Println(i, ok)

	for i := 0; i < 100; i++ {
		hash_table[i] = "dfgdgdfg"
	}

	// for i, _ := range hash_table {
	// 	fmt.Println("Key:", i)
	// }
	//fmt.Println(len(hash_table))
	return hash_table
}

// #hash
func (hs Hash) longestSubSequence(list []int) {

	m := make(map[int]int)
	n := len(list)
	m[list[0]] = 1
	for i := 1; i < n; i++ {
		x := list[i]
		if m[x-1] > 0 {
			m[x] = m[x-1] + 1
		} else {
			m[x] = 1
		}
	}

	fmt.Println(m)

}

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)

	if k > n {
		return res
	}

	hasthbl := make(map[string]bool)
	queue := make([]string, 0)

	for i := 1; i <= 9; i++ {
		v := strconv.Itoa(i)
		queue = append(queue, v)
	}

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if len(head) == k {
			tmp := 0
			key := ""
			chars := []byte(head)
			comb := make([]int, 0)
			for i := 0; i < k; i++ {
				v, _ := strconv.Atoi(string(chars[i]))
				comb = append(comb, v)
				tmp += v
			}
			if tmp == n {
				sort.Slice(chars, func(i, j int) bool { return string(chars[i]) > string(chars[j]) })
				key = string(chars)
				if hasthbl[key] == false {
					hasthbl[key] = true
					res = append(res, comb)
				}
			}

		} else {
			for i := 1; i <= 9; i++ {
				v := strconv.Itoa(i)
				if !strings.Contains(head, v) {
					queue = append(queue, head+v)
				}
			}
		}

	}

	return res
}

func MakeNewHashMap() map[int][]int {
	hm := make(map[int][]int)

	for i := 0; i < 10; i++ {
		hm[i] = make([]int, 0)
	}

	return hm
}
