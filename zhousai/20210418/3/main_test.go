package main

import (
	"reflect"
	"testing"
)

func Test_getOrder(t *testing.T) {
	type args struct {
		tasks [][]int
	}
	tasks := [][]int{}
	tasks = append(tasks,[]int{1,2})
	tasks = append(tasks,[]int{2,4})
	tasks = append(tasks,[]int{3,2})
	tasks = append(tasks,[]int{4,1})
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				tasks,
			},
			[]int{0,2,3,1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOrder(tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
