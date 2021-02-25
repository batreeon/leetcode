package main
import "testing"
func TestFindMedianSortedArrays(t *testing.T) {
	got := findMedianSortedArrays([]int{1,3},[]int{2})
	want := 2.0
	if got != want {
		t.Error("got",got,"want",want)
	}
}