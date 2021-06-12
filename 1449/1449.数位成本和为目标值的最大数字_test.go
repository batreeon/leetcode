/*
 * @lc app=leetcode.cn id=1449 lang=golang
 *
 * [1449] 数位成本和为目标值的最大数字
 */

// @lc code=start

package main

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		cost   []int
		target int
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
		// 		[]int{4,3,2,5,6,7,2,5,5},
		// 		9,
		// 	},
		// 	"7772",
		// },
		// {
		// 	"2",
		// 	args{
		// 		[]int{5,6,7,3,4,6,7,4,8},
		// 		29,
		// 	},
		// 	"884444444",
		// },
		// {
		// 	"2",
		// 	args{
		// 		[]int{70,84,55,63,74,44,27,76,34},
		// 		659,
		// 	},
		// 	"99977777777777777777776",
		// },
		{
			"2",
			args{
				[]int{4,2,3,3,4,5,5,2,2},
				795,
			},
			"9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999994",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.cost, tt.args.target); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
