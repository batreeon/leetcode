/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */

// @lc code=start

package main

import "testing"

func Test_judgeSquareSum(t *testing.T) {
	type args struct {
		c int
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
				10000000,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeSquareSum(tt.args.c); got != tt.want {
				t.Errorf("judgeSquareSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
