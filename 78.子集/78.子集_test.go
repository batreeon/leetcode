package main
import "testing"
func TestSubsets(t *testing.T) {
	got := subsets([]int{9,0,3,5,7})
	t.Error("got:\n",got)
}