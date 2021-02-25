package main
import "testing"
func TestSearch(t *testing.T) {
	// got := search([]int{4,5,6,7,0,1,2},0)
	// want := 4
	got := search([]int{5,1,3},0)
	want := -1
	if got != want {
		t.Error("got",got,"want",want)
	}
}