/*
 * @lc app=leetcode.cn id=881 lang=golang
 *
 * [881] 救生艇
 */

// @lc code=start

package main

import "testing"

func Test_numRescueBoats(t *testing.T) {
	type args struct {
		people []int
		limit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				[]int{2,49,10,7,11,41,47,2,22,6,13,12,33,18,10,26,2,6,50,10},
				50,
			},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRescueBoats(tt.args.people, tt.args.limit); got != tt.want {
				t.Errorf("numRescueBoats() = %v, want %v", got, tt.want)
			}
		})
	}
}
