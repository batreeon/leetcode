/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start

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
		// TODO: Add test cases.
		// {
		// 	"1",
		// 	args{
		// 		string("3[a]2[bc]"),
		// 	},
		// 	string("aaabcbc"),
		// },
		// {
		// 	"1",
		// 	args{
		// 		string("3[a2[c]]"),
		// 	},
		// 	string("accaccacc"),
		// },
		// {
		// 	"1",
		// 	args{
		// 		string("2[abc]3[cd]ef"),
		// 	},
		// 	string("abcabccdcdcdef"),
		// },
		// {
		// 	"1",
		// 	args{
		// 		string("abc3[cd]xyz"),
		// 	},
		// 	string("abccdcdcdxyz"),
		// },
		{
			"1",
			args{
				string("3[z]2[2[y]pq4[2[jk]e1[f]]]ef"),
			},
			string("zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef"),
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
