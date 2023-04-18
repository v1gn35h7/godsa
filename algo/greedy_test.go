package algo

import (
	"testing"
)

func TestFractionalKanpSack(t *testing.T) {
	list := make([]KnapsackItem, 0)
	list = append(list, KnapsackItem{P: 60, W: 10})
	list = append(list, KnapsackItem{P: 100, W: 20})
	list = append(list, KnapsackItem{P: 120, W: 30})
	want := 240
	got := FractionalKanpsack(50, list)

	if want != got {
		t.Errorf("Failed... got: %d want: %d", got, want)
	}
}

func BenchmarkFractionalSack(b *testing.B) {
	list := make([]KnapsackItem, 0)
	list = append(list, KnapsackItem{P: 60, W: 10})
	list = append(list, KnapsackItem{P: 100, W: 20})
	list = append(list, KnapsackItem{P: 120, W: 30})
	for i := 0; i < b.N; i++ {
		FractionalKanpsack(50, list)
	}
}
