package main

import "testing"

func Test_numDifferentIntegers(t *testing.T) {
	type args struct {
		word string
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
				"uu717761265362523668772526716127260222200144937319826j717761265362523668772526716127260222200144937319826k717761265362523668772526716127260222200144937319826b7177612653625236687725267161272602222001449373198262hgb9utohfvfrxprqva3y5cdfdudf7zuh64mobfr6mhy17",
			},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDifferentIntegers(tt.args.word); got != tt.want {
				t.Errorf("numDifferentIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
