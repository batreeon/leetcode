/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 */

// @lc code=start
package main
func openLock(deadends []string, target string) int {
	deads := map[string]bool{}
	for _,deadend := range deadends {
		deads[deadend] = true
	}
	if deads[string("0000")] {
		return -1
	}
	q := []string{string("0000")}
	visited := map[string]bool{string("0000"):true}

	var add_one func(string,int) string
	var sub_one func(string,int) string
	add_one = func(x string,i int) string {
		xb := []byte(x)
		xbi_int := xb[i]-'0'
		xbi_int = (xbi_int+1)%10
		xb[i] = '0' + xbi_int
		return string(xb)
	}
	sub_one = func(x string,i int) string {
		xb := []byte(x)
		xbi_int := xb[i]-'0'
		// 实际上不用该用%这个取余操作，应该用取模操作
		xbi_int = (xbi_int-1+10)%10
		xb[i] = '0' + xbi_int
		return string(xb)
	}

	step := 0
	for len(q) > 0 {
		l := len(q)
		for i := range q {
			x := q[i]
			if x == target {
				return step
			}
			if deads[x] {
				continue
			}
			for j := 0 ; j < 4 ; j++ {
				y := add_one(x,j)
				if !visited[y] {
					q = append(q,y)
					visited[y] = true
				}
				z := sub_one(x,j)
				if !visited[z] {
					q = append(q,z)
					visited[z] = true
				}
			}
		}
		q = q[l:]
		step++
	}


	return -1
}
// @lc code=end

