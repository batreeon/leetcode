/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start

package main

import "testing"

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}
	args1 := make([][]byte,5)
	args1[0] = []byte{'1','1','1','1','0'}
	args1[1] = []byte{'1','1','1','1','0'}
	args1[2] = []byte{'1','1','1','1','1'}
	args1[3] = []byte{'1','1','1','1','1'}
	args1[4] = []byte{'0','0','1','1','1'}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				args1,
			},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
