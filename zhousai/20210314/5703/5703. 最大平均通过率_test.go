/*
一所学校里有一些班级，每个班级里有一些学生，现在每个班都会进行一场期末考试。给你一个二维数组 classes ，其中 classes[i] = [passi, totali] ，表示你提前知道了第 i 个班级总共有 totali 个学生，其中只有 passi 个学生可以通过考试。

给你一个整数 extraStudents ，表示额外有 extraStudents 个聪明的学生，他们 一定 能通过任何班级的期末考。你需要给这 extraStudents 个学生每人都安排一个班级，使得 所有 班级的 平均 通过率 最大 。

一个班级的 通过率 等于这个班级通过考试的学生人数除以这个班级的总人数。平均通过率 是所有班级的通过率之和除以班级数目。

请你返回在安排这 extraStudents 个学生去对应班级后的 最大 平均通过率。与标准答案误差范围在 10-5 以内的结果都会视为正确结果。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-average-pass-ratio
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "testing"

func Test_maxAverageRatio(t *testing.T) {
	type args struct {
		classes       [][]int
		extraStudents int
	}
	classes := [][]int{}
	classes = append(classes, []int{2,4})
	classes = append(classes, []int{3,9})
	classes = append(classes, []int{4,5})
	classes = append(classes, []int{2,10})
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				classes,
				4,
			},
			0.53485,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAverageRatio(tt.args.classes, tt.args.extraStudents); got != tt.want {
				t.Errorf("maxAverageRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}
