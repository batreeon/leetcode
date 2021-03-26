/*
 * @lc app=leetcode.cn id=1209 lang=golang
 *
 * [1209] 删除字符串中的所有相邻重复项 II
 */

// @lc code=start

package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				string("deeedbbcccbdaa"),
				3,
			},
			string("aa"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
