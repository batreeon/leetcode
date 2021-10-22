/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 */

// @lc code=start

package main

import "testing"

func Test_isScramble(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		// {
		// 	"1",
		// 	args{
		// 		"great",
		// 		"rgeat",
		// 	},
		// 	true,
		// },
		// {
		// 	"2",
		// 	args{
		// 		"abcdbdacbdac",
		// 		"bdacabcdbdac",
		// 	},
		// 	true,
		// },
		// {
		// 	"3",
		// 	args{
		// 		"abb",
		// 		"bba",
		// 	},
		// 	true,
		// },
		{
			"4",
			args{
				"eebaacbcbcadaaedceaaacadccd",
				"eadcaacabaddaceacbceaabeccd",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isScramble(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isScramble() = %v, want %v", got, tt.want)
			}
		})
	}
}
