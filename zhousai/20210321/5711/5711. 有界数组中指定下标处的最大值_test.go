package main

import "testing"

func Test_maxValue(t *testing.T) {
	type args struct {
		n      int
		index  int
		maxSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	"1",
		// 	args{
		// 		4,2,6,
		// 	},
		// 	2,
		// },
		// {
		// 	"2",
		// 	args{
		// 		6,1,10,
		// 	},
		// 	3,
		// },
		{
			"3",
			args{
				3,2,18,
			},
			7,
		},
		{
			"4",
			args{
				685453290,
				293811406,
				689728311,
			},
			-1,// 未知
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.n, tt.args.index, tt.args.maxSum); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
