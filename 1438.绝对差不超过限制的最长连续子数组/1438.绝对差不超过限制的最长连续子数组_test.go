package main 
import "testing"
func TestLongestSubarray(t *testing.T) {
	ans := longestSubarray([]int{10,1,2,4,7,2},5)
	if ans != 4 {
		t.Error("ans",ans,",want",4)
	}
}