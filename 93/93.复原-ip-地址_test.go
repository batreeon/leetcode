/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start

package main

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				string("0000"),
			},
			[]string{"0.0.0.0"},
		},
		{
			"2",
			args{
				string("25525511135"),
			},
			[]string{"255.255.11.135","255.255.111.35"},
		},
		{
			"3",
			args{
				string("010010"),
			},
			[]string{"0.10.0.10","0.100.1.0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
