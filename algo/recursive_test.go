package algo

import "testing"

func TestAllSubSets(t *testing.T) {
	got := PrintAllSubsets("123")
	want := map[string]bool{"123": true, "12": true, "13": true, "1": true, "2": true, "3": true}
	if len(got) != len(want) {
		t.Fail()
	}
}
