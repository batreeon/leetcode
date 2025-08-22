/*
 * @lc app=leetcode.cn id=885 lang=golang
 *
 * [885] 螺旋矩阵 III
 */

// @lc code=start
package main

import (
	"reflect"
	"testing"
)

func Test_spiralMatrixIII(t *testing.T) {
	type args struct {
		rows   int
		cols   int
		rStart int
		cStart int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"0",
			args{
				5,6,1,4,
			},
			[][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralMatrixIII(tt.args.rows, tt.args.cols, tt.args.rStart, tt.args.cStart); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralMatrixIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
