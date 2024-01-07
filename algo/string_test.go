package algo

import "testing"

func TestSubstr(t *testing.T) {
	want := true

	got := substr("go", "gohan")

	if want != got {
		t.Errorf("Test Failed")
	}

}

func BenchmarkSubstr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = substr("go", "gohan")
	}
}
