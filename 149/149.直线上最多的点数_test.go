/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] 直线上最多的点数
 */

// @lc code=start

package main

import "testing"

func Test_maxPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	ps := [][]int{}
	ps = append(ps,[]int{1,1})
	ps = append(ps,[]int{2,2})
	ps = append(ps,[]int{3,3})
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				ps,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
