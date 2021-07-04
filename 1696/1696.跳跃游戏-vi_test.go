/*
 * @lc app=leetcode.cn id=1696 lang=golang
 *
 * [1696] 跳跃游戏 VI
 */

// @lc code=start

package main

import "testing"

func Test_maxResult(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				[]int{1,-1,-2,4,-7,3},
				2,
			},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxResult(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
