package ds

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// func main() {
// 	ioreader := bufio.NewReader(os.Stdin)
// 	inp := scanLine(ioreader)
// 	difs := strings.Split(inp, ", ")
// 	//printPriorityQueue(difs)
// 	//kthLargestElemenet(difs, 3)
// 	//kLargestElemenets(difs, 5)
// 	sortKSortedArray(difs, 3)
// }

// Finf K-Largest elememnts
func kLargestElemenets(difs []string, k int) {
	intArry := make(IntHeap, 0)
	for _, d := range difs {
		tmp, _ := strconv.Atoi(strings.TrimSpace(d))
		intArry.Push(tmp)
	}

	heap.Init(&intArry)

	for i := 0; i < k; i++ {
		fmt.Printf("%d ", heap.Pop(&intArry))
	}
	fmt.Printf("\n")
}

// Find Kthe Largest/Smallest elememnt
func kthLargestElemenet(difs []string, k int) {
	intArry := make(IntHeap, 0)
	heap.Init(&intArry)
	for _, d := range difs {
		tmp, _ := strconv.Atoi(strings.TrimSpace(d))
		heap.Push(&intArry, tmp)

		if intArry.Len() > k {
			heap.Pop(&intArry)
		}
	}

	fmt.Println(heap.Pop(&intArry))
}

func printPriorityQueue(difs []string) {
	intArry := make(IntHeap, 0)
	heap.Init(&intArry)
	for _, d := range difs {
		tmp, _ := strconv.Atoi(strings.TrimSpace(d))
		heap.Push(&intArry, tmp)
	}

	fmt.Println("Priority Queue: ", intArry)
}

func sortKSortedArray(difs []string, k int) {
	intArry := make(IntHeap, 0)
	heap.Init(&intArry)
	for i, d := range difs {
		if i > k+1 {
			break
		}
		tmp, _ := strconv.Atoi(strings.TrimSpace(d))
		heap.Push(&intArry, tmp)
	}

	for j := k + 2; j < len(difs); j++ {
		fmt.Printf("%d ", heap.Pop(&intArry))
		tmp, _ := strconv.Atoi(strings.TrimSpace(difs[j]))
		heap.Push(&intArry, tmp)
	}

	for intArry.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(&intArry))
	}
}
