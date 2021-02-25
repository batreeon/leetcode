package main
import "testing"
func TestUniquePaths(t *testing.T) {
	// got := uniquePaths(2,3)
	// want := 3
	got := uniquePaths(23,12)
	want := 193536720
	if got != want {
		t.Error("got",got,"want",want)
	}
}