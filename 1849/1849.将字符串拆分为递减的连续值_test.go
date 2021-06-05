/*
 * @lc app=leetcode.cn id=1849 lang=golang
 *
 * [1849] 将字符串拆分为递减的连续值
 */

// @lc code=start

package main

import "testing"

func Test_splitString(t *testing.T) {
	type args struct {
		s string
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
				string("050043"),
			},
			true,
		},
		{
			"2",
			args{
				string("4771447713"),
			},
			true,
		},
		{
			"3",
			args{
				string("0166537080"),
			},
			false,
		},
		{
			"4",
			args{
				string("10"),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitString(tt.args.s); got != tt.want {
				t.Errorf("splitString() = %v, want %v", got, tt.want)
			}
		})
	}
}
