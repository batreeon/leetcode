/*
 * @lc app=leetcode.cn id=735 lang=golang
 *
 * [735] 行星碰撞
 */

// @lc code=start

package main

import (
	"reflect"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
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
		// 		[]int{10,2,-5},
		// 	},
		// 	[]int{10},
		// },
		// {
		// 	"2",
		// 	args{
		// 		[]int{-2,-1,1,2},
		// 	},
		// 	[]int{-2,-1,1,2},
		// },
		{
			"3",
			args{
				[]int{1,-2,-2,-2},
			},
			[]int{-2,-2,-2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidCollision(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
