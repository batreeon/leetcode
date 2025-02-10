/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */

// @lc code=start
// 广度优先遍历
func minMutation(startGene string, endGene string, bank []string) int {
	q := []string{startGene}
	visited := map[string]struct{}{startGene: struct{}{}}
	res := 0
	for len(q) > 0 {
		l := len(q)
		for _, gene := range q[:l] {
			for _, b := range bank {
				if _, ok := visited[b]; ok {
					continue
				}
				if near(gene, b) {
					if b == endGene {
						return res+1
					}
					q = append(q, b)
					visited[b] = struct{}{}
				}
			}
		}
		res++
		q = q[l:]
	}
	return -1
}

func near(startGene, endGene string) bool {
	if startGene == endGene {
		return false
	}
	changeNum := 0
	for i := 0; i < 8; i++ {
		if startGene[i] != endGene[i] {
			changeNum++
			if changeNum > 1 {
				return false
			}
		}
	}
	return true
}
// @lc code=end

