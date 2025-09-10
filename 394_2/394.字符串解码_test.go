/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
// 2[abc]3[cd]ef
// stack
// 2
// 2[
// 3[2[abc]] => 3[abcabc]
package main

import "testing"

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"0",
			args{
				"3[a]2[bc]",
			},
			"aaabcbc",
		},
		{
			name: "0",
			args: args{
				"3[a2[c]]",
			},
			want: "accaccacc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
