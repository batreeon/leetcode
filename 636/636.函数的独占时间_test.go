/*
 * @lc app=leetcode.cn id=636 lang=golang
 *
 * [636] 函数的独占时间
 */

// @lc code=start

package main

import (
	"reflect"
	"testing"
)

func Test_exclusiveTime(t *testing.T) {
	type args struct {
		n    int
		logs []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				2,
				[]string{
					"0:start:0","1:start:2","1:end:5","0:end:6",
				},
			},
			[]int{3,4},
		},
		// {
		// 	"2",
		// 	args{
		// 		1,
		// 		[]string{
		// 			"0:start:0","0:start:2","0:end:5","0:start:6","0:end:6","0:end:7",
		// 		},
		// 	},
		// 	[]int{8},
		// },
		// {
		// 	"3",
		// 	args{
		// 		2,
		// 		[]string{
		// 			"0:start:0","0:start:2","0:end:5","1:start:6","1:end:6","0:end:7",
		// 		},
		// 	},
		// 	[]int{7,1},
		// },
		// {
		// 	"4",
		// 	args{
		// 		1,
		// 		[]string{
		// 			"0:start:0","0:start:2","0:end:5","0:start:6","0:end:6","0:end:7",
		// 		},
		// 	},
		// 	[]int{8},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exclusiveTime(tt.args.n, tt.args.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exclusiveTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
