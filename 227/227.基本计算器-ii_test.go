/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start

package main

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
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
		// 		string("3+2*2"),
		// 	},
		// 	7,
		// },
		// {
		// 	"1",
		// 	args{
		// 		string("1+1+1"),
		// 	},
		// 	3,
		// },
		{
			"1",
			args{
				string("1*2-3/4+5*6-7*8+9/10"),
			},
			-24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
