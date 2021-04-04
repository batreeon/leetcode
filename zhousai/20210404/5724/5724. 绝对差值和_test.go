package main

import "testing"

func Test_minAbsoluteSumDiff(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				[]int{1,10,4,4,2,7},
				[]int{9,3,5,1,7,4},
			},
			20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAbsoluteSumDiff(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minAbsoluteSumDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
