/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 */

// @lc code=start

package main

import "testing"

func Test_openLock(t *testing.T) {
	type args struct {
		deadends []string
		target   string
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
				[]string{"8888"},
				string("0009"),
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openLock(tt.args.deadends, tt.args.target); got != tt.want {
				t.Errorf("openLock() = %v, want %v", got, tt.want)
			}
		})
	}
}
