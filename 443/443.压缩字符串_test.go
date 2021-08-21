/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 */

// @lc code=start

package main

import "testing"

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
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
				[]byte(string("aabbccc")),
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}
