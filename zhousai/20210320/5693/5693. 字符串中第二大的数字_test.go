package main

import "testing"

func Test_secondHighest(t *testing.T) {
	type args struct {
		s string
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
				string("dfa12321afd"),
			},
			2,
		},
		{
			"2",
			args{
				string("abc1111"),
			},
			-1,
		},
		{
			"3",
			args{
				string("xyz"),
			},
			-1,
		},
		{
			"4",
			args{
				string("ck077"),
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondHighest(tt.args.s); got != tt.want {
				t.Errorf("secondHighest() = %v, want %v", got, tt.want)
			}
		})
	}
}
