/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 */

// @lc code=start

package main

import "testing"

func Test_getHint(t *testing.T) {
	type args struct {
		secret string
		guess  string
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
				"1123",
				"0111",
			},
			"1A1B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint(tt.args.secret, tt.args.guess); got != tt.want {
				t.Errorf("getHint() = %v, want %v", got, tt.want)
			}
		})
	}
}
