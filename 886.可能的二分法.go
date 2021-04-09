/*
 * @lc app=leetcode.cn id=886 lang=golang
 *
 * [886] 可能的二分法
 */

// @lc code=start
func possibleBipartition(N int, dislikes [][]int) bool {
	color := make([]int,N+1)
	uncolor,red,green := 0,1,2

	dislikess := make([][]int,N+1)
	for _,dislike := range dislikes {
		dislikess[dislike[0]] = append(dislikess[dislike[0]],dislike[1])
		dislikess[dislike[1]] = append(dislikess[dislike[1]],dislike[0])
	}

	for i := 1 ; i <= N ; i++ {
		if color[i] == uncolor {
			color[i] = red
			q := []int{i}

			for len(q) > 0 {
				v := q[0]
				q = q[1:]
				nextcolor := red
				if color[v] == red {
					nextcolor = green
				}
				for _,next := range dislikess[v] {
					if color[next] == uncolor {
						color[next] = nextcolor
						q = append(q,next)
					}else if color[next] != nextcolor {
						return false
					}
				}
			}
		}
	}
	
	return true
}
// @lc code=end

