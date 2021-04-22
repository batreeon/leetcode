/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start

package main

import (
	"reflect"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}
	ans := [][]int{}
	ans = append(ans,[]int{1,2,4})
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		// {
		// 	"1",
		// 	args{
		// 		3,7,
		// 	},
		// 	ans,
		// },
		{
			"1",
			args{
				3,9,
			},
			ans,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.args.k, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
