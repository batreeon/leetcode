package main

import "testing"

func Test_numberOfWeakCharacters(t *testing.T) {
	type args struct {
		properties [][]int
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
		// 		[][]int{{7,9},{10,7},{6,9},{10,4},{7,5},{7,10}},
		// 	},
		// 	2,
		// },
		{
			"1",
			args{
				[][]int{{1,5},{10,4},{4,3}},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWeakCharacters(tt.args.properties); got != tt.want {
				t.Errorf("numberOfWeakCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
