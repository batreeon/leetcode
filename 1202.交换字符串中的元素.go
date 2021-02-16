/*
 * @lc app=leetcode.cn id=1202 lang=golang
 *
 * [1202] 交换字符串中的元素
 */

// @lc code=start
type sb []byte
func (b sb) Len() int { return len(b) }
func (b sb) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b sb) Less(i, j int) bool { return b[i] < b[j] }
func smallestStringWithSwaps(s string, pairs [][]int) string {
	parent,size := make([]int,len(s)),make([]int,len(s))
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	var find func(x int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x,y int) bool {
		fx,fy := find(x),find(y)
		if fx == fy {
			return false
		}
		if size[fx] < size[fy] {
			fx,fy = fy,fx
		}
		size[fx] += size[fy]
		parent[fy] = fx
		return true
	}
	for _,v := range pairs {
		union(v[0],v[1])
	}
	i2parent := map[int]sb{}
	for i := range s {
		ip := find(i)
		i2parent[ip] = append(i2parent[ip],s[i])
		sort.Sort(i2parent[ip])
	}
	var ans sb
	for i := range s {
		ans = append(ans,i2parent[find(i)][0])
		i2parent[find(i)] = i2parent[find(i)][1:]
	}
	return string(ans)
}
// @lc code=end

