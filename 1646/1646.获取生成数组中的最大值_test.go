/*
 * @lc app=leetcode.cn id=1646 lang=golang
 *
 * [1646] 获取生成数组中的最大值
 */

// @lc code=start

package main

import "testing"

func Test_getMaximumGenerated(t *testing.T) {
	type args struct {
		n int
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
				9,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumGenerated(tt.args.n); got != tt.want {
				t.Errorf("getMaximumGenerated() = %v, want %v", got, tt.want)
			}
		})
	}
}
