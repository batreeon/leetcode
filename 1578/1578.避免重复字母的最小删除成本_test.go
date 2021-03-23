/*
 * @lc app=leetcode.cn id=1578 lang=golang
 *
 * [1578] 避免重复字母的最小删除成本
 */

// @lc code=start

package main

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		s    string
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	"1",
		// 	args{
		// 		string("abaac"),
		// 		[]int{1,2,3,4,5},
		// 	},
		// 	3,
		// },
		// {
		// 	"2",
		// 	args{
		// 		string("aabaa"),
		// 		[]int{1,2,3,4,1},
		// 	},
		// 	2,
		// },
		// {
		// 	"3",
		// 	args{
		// 		string("aaabbbabbbb"),
		// 		[]int{3,5,10,7,5,3,5,5,4,8,1},
		// 	},
		// 	26,
		// },
		{
			"4",
			args{
				string(""),
				[]int{3,5,10,7,5,3,5,5,4,8,1},
			},
			26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.s, tt.args.cost); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
