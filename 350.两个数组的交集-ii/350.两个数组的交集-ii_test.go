package main
import "testing"
func TestIntersect(t *testing.T) {
	got := intersect([]int{1,2,2,1},[]int{2,2})
	t.Error("got",got,"want","[2,2]")
}