package ds

import "fmt"

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
