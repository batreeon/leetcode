/*
 * @lc app=leetcode.cn id=1718 lang=golang
 *
 * [1718] 构建字典序最大的可行序列
 */

// @lc code=start

package main

import (
	"reflect"
	"testing"
)

func Test_constructDistancedSequence(t *testing.T) {
	type args struct {
		n int
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
		// 		3,
		// 	},
		// 	[]int{3,1,2,3,2},
		// },
		// {
		// 	"2",
		// 	args{
		// 		4,
		// 	},
		// 	[]int{4,2,3,2,4,3,1},
		// },
		{
			"3",
			args{
				5,
			},
			[]int{5,3,1,4,3,5,2,4,2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructDistancedSequence(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructDistancedSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
