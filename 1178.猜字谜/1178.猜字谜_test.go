package main
import "testing"
func TestFindNumOfValidWords(t *testing.T) {
	ws := []string{"aaaa","asas","able","ability","actt","actor","access"}
	ps := []string{"aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"}
	got := findNumOfValidWords(ws,ps)
	want := []int{1,1,3,2,4,0}
	t.Error("got",got,"\nwant",want)
}