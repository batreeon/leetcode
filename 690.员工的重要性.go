/*
 * @lc app=leetcode.cn id=690 lang=golang
 *
 * [690] 员工的重要性
 */

// @lc code=start
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
	importances := map[int]int{}
    nexts := map[int][]int{}
	for _,employee := range employees {
		importances[employee.Id] = employee.Importance
		nexts[employee.Id] = []int{}
		for _,next := range employee.Subordinates {
			nexts[employee.Id] = append(nexts[employee.Id],next)
		}
	}
	result := 0
	allnext := []int{id}
	for len(allnext) > 0 {
		l := len(allnext)
		for i := 0 ; i < l ; i++ {
			result += importances[allnext[i]]
			allnext = append(allnext,nexts[allnext[i]]...)
		}
		allnext = allnext[l:]
	}

	return result
}
// @lc code=end

