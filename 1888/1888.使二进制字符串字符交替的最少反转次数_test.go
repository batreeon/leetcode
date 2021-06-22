/*
 * @lc app=leetcode.cn id=1888 lang=golang
 *
 * [1888] 使二进制字符串字符交替的最少反转次数
 */

// @lc code=start

package main

import "testing"

func Test_minFlips(t *testing.T) {
	type args struct {
		s string
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
				string("010"),
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlips(tt.args.s); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
