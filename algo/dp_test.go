package algo

import "testing"

func TestFibonacci(t *testing.T) {
	want := 34
	got := Fibonacci(9)

	if want != got {
		t.Errorf("Ailed got: %d -> want: %d", got, want)
	}

}
