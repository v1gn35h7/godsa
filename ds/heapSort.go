package ds

import "fmt"

type Heap struct{}

func GetNewHeap() {
	// ioreader := bufio.NewReader(os.Stdin)
	// inp := scanLine(ioreader)
	// difs := strings.Split(inp, ", ")
	// intArry := make([]int, 0)
	// for _, d := range difs {
	// 	tmp, _ := strconv.Atoi(strings.TrimSpace(d))
	// 	intArry = append(intArry, tmp)
	// }

	// nonRoot := (len(intArry) / 2) - 1

	// for i := nonRoot; i >= 0; i-- {
	// 	intArry = heapify(intArry, len(intArry), i)
	// }

	// x := len(intArry)

	// for i := x; i > 0; i-- {
	// 	intArry[0], intArry[i-1] = intArry[i-1], intArry[0]
	// 	intArry = heapify(intArry, i-1, 0)
	// }

	// fmt.Println(intArry)
}

func (heap *Heap) heapify(arr []int, n, i int) []int {
	largest := i
	left := (i * 2) + 1
	right := (i * 2) + 2

	if left < n && arr[largest] < arr[left] {
		largest = left
	}

	if right < n && arr[largest] < arr[right] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		return heap.heapify(arr, n, largest)
	}
	return arr
}

// {1, 0, 1, 1, 1, 0, 0}
// {1, 1, 1, 1}
func equal_o_1_sub_array(list []int) {

	var max_len, end_index, sum int
	hm := make(map[int]int)
	// n := len(list)

	for i, v := range list {
		if v == 0 {
			list[i] = -1
		}
	}

	//fmt.Println(list)

	for i, _ := range list {

		sum += list[i]

		if sum == 0 {
			max_len = i + 1
			end_index = i
		}

		if _, ok := hm[sum]; ok {
			if max_len < i-(hm[sum]+1) {
				max_len = i - (hm[sum] + 1)
				end_index = i
			}
		} else {
			hm[sum] = i
		}
	}

	fmt.Println(list, hm, max_len, end_index)

	start := end_index - max_len + 1

	fmt.Println(start, " to ", end_index)
}
