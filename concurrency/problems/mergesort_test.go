package problems

import (
	"runtime"
	"slices"
	"testing"
)

func TestMergeSort(t *testing.T) {
	want := []int64{0, 2, 3, 3, 3, 4, 4, 6, 6, 6, 7, 7, 32, 34, 35, 75, 88, 234, 7858}

	res := MergeSort([]int64{4, 6, 3, 7, 4, 6, 7, 3, 34, 7858, 234, 75, 0, 32, 35, 6, 2, 88, 3})
	got := <-res
	if ok := slices.Equal(want, got); ok != true {
		t.Errorf("Concurent merge sorting Failed ")
	}
}

func BenchmarkMergeSort(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < b.N; i++ {
		res := MergeSort([]int64{4, 6, 3, 7, 4, 6, 7, 3, 34, 7858, 234, 75, 0, 32, 35, 6, 2, 88, 3})
		<-res
	}
}
