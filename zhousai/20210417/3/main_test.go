package main

import (
	"reflect"
	"testing"
)

func Test_getMaximumXor(t *testing.T) {
	type args struct {
		nums       []int
		maximumBit int
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
				[]int{0,1,1,3},
				2,
			},
			[]int{0,3,2,3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumXor(tt.args.nums, tt.args.maximumBit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMaximumXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
