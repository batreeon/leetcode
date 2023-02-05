/*
 * @lc app=leetcode.cn id=1210 lang=golang
 *
 * [1210] 穿过迷宫的最少移动次数
 */

// @lc code=start
package main

import "testing"

func Test_minimumMoves(t *testing.T) {
	type args struct {
		grid [][]int
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
				grid: [][]int{
					{0,0,0,0,0,1},
					{1,1,0,0,1,0},
					{0,0,0,0,1,1},
					{0,0,1,0,1,0},
					{0,1,1,0,0,0},
					{0,1,1,0,0,0},
				},
			},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumMoves(tt.args.grid); got != tt.want {
				t.Errorf("minimumMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
