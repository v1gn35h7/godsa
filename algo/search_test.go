package algo

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	inp := []string{"1", "4", "6", "8", "10"}
	want := 1
	got := binary_search(inp, 0, len(inp)-1, 4)

	if got != want {
		t.Errorf("Got %d inseatd fof %d", got, want)
	}

}

func BenchmarkBinarySearch(b *testing.B) {
	inp := []string{"1", "4", "6", "8", "10"}
	for i := 0; i < b.N; i++ {
		binary_search(inp, 0, len(inp)-1, 10)
	}
}

func ExampleBinarySearch() {
	inp := []string{"1", "4", "6", "8", "10"}
	fmt.Println(binary_search(inp, 0, len(inp)-1, 10))
	// 4
}
