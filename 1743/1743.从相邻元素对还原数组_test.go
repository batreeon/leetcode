/*
 * @lc app=leetcode.cn id=1743 lang=golang
 *
 * [1743] 从相邻元素对还原数组
 */

// @lc code=start

package main

import (
	"reflect"
	"testing"
)

func Test_restoreArray(t *testing.T) {
	type args struct {
		adjacentPairs [][]int
	}
	adjacentPairs := [][]int{}
	adjacentPairs = append(adjacentPairs,[]int{2,1})
	adjacentPairs = append(adjacentPairs,[]int{3,4})
	adjacentPairs = append(adjacentPairs,[]int{3,2})
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				adjacentPairs,
			},
			[]int{1,2,3,4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreArray(tt.args.adjacentPairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
