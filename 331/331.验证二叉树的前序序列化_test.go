/*
 * @lc app=leetcode.cn id=331 lang=golang
 *
 * [331] 验证二叉树的前序序列化
 */

// @lc code=start

package main

import "testing"

func Test_isValidSerialization(t *testing.T) {
	type args struct {
		preorder string
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
				"9,3,4,#,#,1,#,#,2,#,6,#,#",
			},
			true,
		},

		{
			"2",
			args{
				"1,#,#,1",
			},
			false,
		},

		{
			"3",
			args{
				"1,#,#,#,#",
			},
			false,
		},

		{
			"4",
			args{
				"#,#,#",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSerialization(tt.args.preorder); got != tt.want {
				t.Errorf("isValidSerialization() = %v, want %v", got, tt.want)
			}
		})
	}
}
