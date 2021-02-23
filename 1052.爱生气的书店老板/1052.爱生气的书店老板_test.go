package main
import "testing"
func TestMaxSatisfied(t *testing.T) {
	got := maxSatisfied([]int{1,0,1,2,1,1,7,5},[]int{0,1,0,1,0,1,0,1},3)
	if got != 16 {
		t.Error("got",got,"want",16)
	}
}