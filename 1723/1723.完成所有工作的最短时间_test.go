/*
 * @lc app=leetcode.cn id=1723 lang=golang
 *
 * [1723] 完成所有工作的最短时间
 */

// @lc code=start

package main

import "testing"

func Test_minimumTimeRequired(t *testing.T) {
	type args struct {
		jobs []int
		k    int
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
		// 		[]int{3,2,3},
		// 		3,
		// 	},
		// 	3,
		// },
		// {
		// 	"1",
		// 	args{
		// 		[]int{1,3,5,1000},
		// 		4,
		// 	},
		// 	1000,
		// },
		{
			"1",
			args{
				[]int{6518448,8819833,7991995,7454298,2087579,380625,4031400,2905811,4901241,8480231,7750692,3544254},
				4,
			},
			16274131,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTimeRequired(tt.args.jobs, tt.args.k); got != tt.want {
				t.Errorf("minimumTimeRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
