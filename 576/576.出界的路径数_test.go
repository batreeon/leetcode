/*
 * @lc app=leetcode.cn id=576 lang=golang
 *
 * [576] 出界的路径数
 */

// @lc code=start

package main

import "testing"

func Test_findPaths(t *testing.T) {
	type args struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	"1",
		// 	args{
		// 		2,2,2,0,0,
		// 	},
		// 	6,
		// },
		// {
		// 	"2",
		// 	args{
		// 		8,50,23,5,26,
		// 	},
		// 	914783380,
		// },
		{
			"3",
			args{
				36,5,50,15,3,
			},
			390153306,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPaths(tt.args.m, tt.args.n, tt.args.maxMove, tt.args.startRow, tt.args.startColumn); got != tt.want {
				t.Errorf("findPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
