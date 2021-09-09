/*
 * @lc app=leetcode.cn id=68 lang=golang
 *
 * [68] 文本左右对齐
 */

// @lc code=start

package main

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		// {
		// 	"1",
		// 	args{
		// 		[]string{"This", "is", "an", "example", "of", "text", "justification."},
		// 		16,
		// 	},
		// 	[]string{"This    is    an","example  of text","justification.  "},
		// },
		{
			"2",
			args{
				[]string{"What","must","be","acknowledgment","shall","be"},
				16,
			},
			[]string{"What   must   be","acknowledgment  ","shall be        "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
