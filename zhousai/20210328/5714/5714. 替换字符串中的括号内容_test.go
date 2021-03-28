package main

import "testing"

func Test_evaluate(t *testing.T) {
	type args struct {
		s         string
		knowledge [][]string
	}
	kn := [][]string{}
	kn = append(kn,[]string{"name","bob"})
	kn = append(kn,[]string{"age","two"})
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				"(name)is(age)yearsold",
				kn,
			},
			"bobistwoyearsold",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluate(tt.args.s, tt.args.knowledge); got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
