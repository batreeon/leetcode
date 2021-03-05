/*
 * @lc app=leetcode.cn id=735 lang=golang
 *
 * [735] 行星碰撞
 */

// @lc code=start
package main

type stack struct {
	s []int
	l int
}

func (st *stack) push(x int) {
	st.s = append(st.s,x)
	st.l++
}

func (st *stack) top() int {
	return st.s[st.l-1]
}

func (st *stack) pop() {
	st.s = st.s[:st.l-1]
	st.l--
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
func asteroidCollision(asteroids []int) []int {
	st := stack{}
	for _,a := range asteroids {
		if st.l == 0 {
			st.push(a)
		}else{
			top := st.top()
			for top > 0 && a < 0 {
				if abs(top) == abs(a) {
					st.pop()
					break
				}else if abs(top) < abs(a) {
					st.pop()
					if st.l > 0 {
						top = st.top()
					}else{
						st.push(a)
						break
					}
				}else if abs(top) > abs(a) {
					break
				}
			}
			if !(top > 0 && a < 0) {
				st.push(a)
			}
		}
	}
	return st.s
}
// @lc code=end

