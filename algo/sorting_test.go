package algo

import "testing"

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort([]int64{4, 6, 3, 7, 4, 6, 7, 3, 34, 7858, 234, 75, 0, 32, 35, 6, 2, 88, 3})
	}
}
