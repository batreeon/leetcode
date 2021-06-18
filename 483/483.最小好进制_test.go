/*
 * @lc app=leetcode.cn id=483 lang=golang
 *
 * [483] 最小好进制
 */

// @lc code=start

package main

import "testing"

func Test_smallestGoodBase(t *testing.T) {
	type args struct {
		n string
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
				"13",
			},
			"3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestGoodBase(tt.args.n); got != tt.want {
				t.Errorf("smallestGoodBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
