package algo

import "testing"

func TestFibonacci(t *testing.T) {
	want := 34
	got := Fibonacci(9)

	if want != got {
		t.Errorf("Failed got: %d -> want: %d", got, want)
	}

}

func TestKnapSack(t *testing.T) {
	var want int16 = 65
	got := KnapSack01(3, 6, []int16{1, 2, 3}, []int16{10, 15, 40})
	if want != got {
		t.Errorf("Failed got: %d -> want: %d", got, want)
	}
}

func BenchmarkKnapSack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KnapSack01(3, 6, []int16{1, 2, 3}, []int16{10, 15, 40})
	}
}

func TestWordBreak(t *testing.T) {
	want := true

	got := WordBreak("ilike", []string{"i", "like", "sam", "sung", "samsung", "mobile", "ice",
		"cream", "icecream", "man", "go", "mango"})

	if got != want {
		t.Errorf("Test case Failed")
	}
}

func BenchmarkWordBreak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WordBreak("ilike", []string{"i", "like", "sam", "sung", "samsung", "mobile", "ice",
			"cream", "icecream", "man", "go", "mango"})
	}
}
