/*
 * @lc app=leetcode.cn id=600 lang=golang
 *
 * [600] 不含连续1的非负整数
 */

// @lc code=start

package main

import "testing"

func Test_findIntegers(t *testing.T) {
	type args struct {
		n int
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
				16,
			},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIntegers(tt.args.n); got != tt.want {
				t.Errorf("findIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
