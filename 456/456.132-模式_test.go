/*
 * @lc app=leetcode.cn id=456 lang=golang
 *
 * [456] 132模式
 */

// @lc code=start

package main

import "testing"

func Test_find132pattern(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				[]int{1,2,3,4},
			},
			false,
		},
		{
			"2",
			args{
				[]int{3,5,0,3,4},
			},
			true,
		},
		{
			"3",
			args{
				[]int{3,1,4,2},
			},
			true,
		},
		{
			"4",
			args{
				[]int{1,3,2,4,5,6,7,8,9,10},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find132pattern(tt.args.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
