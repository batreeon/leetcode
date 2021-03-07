/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start

package main

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		// {
		// 	"1",
		// 	args{
		// 		[]int{1,1,1,2,2,3},
		// 		2,
		// 	},
		// 	[]int{1,2},
		// },
		// {
		// 	"2",
		// 	args{
		// 		[]int{5,1,-1,-8,-7,8,-5,0,1,10,8,0,-4,3,-1,-1,4,-5,4,-3,0,2,2,2,4,-2,-4,8,-7,-7,2,-8,0,-8,10,8,-8,-2,-9,4,-7,6,6,-1,4,2,8,-3,5,-9,-3,6,-8,-5,5,10,2,-5,-1,-5,1,-3,7,0,8,-2,-3,-1,-5,4,7,-9,0,2,10,4,4,-4,-1,-1,6,-8,-9,-1,9,-9,3,5,1,6,-1,-2,4,2,4,-6,4,4,5,-5},
		// 		7,
		// 	},
		// 	[]int{4,-1,2,-5,-8,0,8},
		// },
		// {
		// 	"1",
		// 	args{
		// 		[]int{1,1,1,2,2,3,4,5,7,8,8,9,4,4,4,5,2,5},
		// 		4,
		// 	},
		// 	[]int{1,2,4,5},
		// },
		// {
		// 	"1",
		// 	args{
		// 		[]int{5,3,1,1,1,3,73,1},
		// 		1,
		// 	},
		// 	[]int{1},
		// },
		// {
		// 	"1",
		// 	args{
		// 		[]int{4,1,-1,2,-1,2,3},
		// 		2,
		// 	},
		// 	[]int{-1,2},
		// },
		{
			"1",
			args{
				[]int{3,0,1,0},
				1,
			},
			[]int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
