/*
 * @lc app=leetcode.cn id=393 lang=golang
 *
 * [393] UTF-8 编码验证
 */

// @lc code=start

package main

import "testing"

func Test_validUtf8(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		// {
		// 	"1",
		// 	args{
		// 		[]int{197,130,1},
		// 	},
		// 	true,
		// },
		// {
		// 	"2",
		// 	args{
		// 		[]int{197},
		// 	},
		// 	false,
		// },
		{
			"3",
			args{
				[]int{240,162,138,147},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validUtf8(tt.args.data); got != tt.want {
				t.Errorf("validUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}
