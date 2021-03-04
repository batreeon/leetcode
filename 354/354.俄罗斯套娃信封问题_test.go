/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start

package main

import (
	"strconv"
	"strings"
	"testing"
)

func Test_maxEnvelopes(t *testing.T) {
	type args struct {
		envelopes [][]int
	}
	// es := [][]int{}
	// es = append(es,[]int{5,6})
	// es = append(es,[]int{6,4})
	// es = append(es,[]int{6,7})
	// es = append(es,[]int{2,3})
	// es := [][]int{}
	// es = append(es,[]int{10,17})
	// es = append(es,[]int{10,19})
	// es = append(es,[]int{16,2})
	// es = append(es,[]int{19,18})
	// es = append(es,[]int{5,6})
	s := string("[30,15],[14,43],[17,18],[35,34],[2,7],[1,12],[11,14],[23,5],[45,13],[18,8],[24,31],[44,31],[47,45],[2,11],[6,44],[3,18],[36,45],[9,28],[30,13],[3,32],[8,20],[27,31],[25,23],[28,5],[5,26],[30,33],[18,46],[15,2],[6,24],[9,4],[5,12],[7,2],[2,27],[23,22]")
	es := [][]int{}
	ss := strings.FieldsFunc(s,func(r rune) bool {
		return r == '[' || r == ',' || r == ']'
	})
	for i,n := range ss {
		num,_ := strconv.Atoi(n)
		if i % 2 == 0 {
			es = append(es,[]int{})
		}
		es[i/2] = append(es[i/2],num)
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{es},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEnvelopes(tt.args.envelopes); got != tt.want {
				t.Errorf("maxEnvelopes() = %v, want %v", got, tt.want)
			}
		})
	}
}
