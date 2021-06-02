/*
 * @lc app=leetcode.cn id=810 lang=golang
 *
 * [810] 黑板异或游戏
 */

// @lc code=start

package main

import "testing"

func Test_xorGame(t *testing.T) {
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
				[]int{0,0,1},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorGame(tt.args.nums); got != tt.want {
				t.Errorf("xorGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
