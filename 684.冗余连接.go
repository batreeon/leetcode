/*
 * @lc app=leetcode.cn id=684 lang=golang
 *
 * [684] 冗余连接
 */

// @lc code=start
type unionFind struct {
	parent, rank []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n+1)
	rank := make([]int, n+1)
	//唉，这里只给了边的个数，我这个是按点来看的，因此根据边数，最多有2n个点
	//n个结点组成的树，有n-1条边，题意是多了一条边，则为n条边。因此根据边数可得点数，使用该数分配数组大小
	//但还要注意，该题的结点标号是　1到n，而不是0到n-1,因此数组大小不能分配为n，要分配为n+1
	for i := range parent {
		parent[i], rank[i] = i, 1
	}
	return &unionFind{parent, rank}
}
func (uf *unionFind) find(x int) int { //找parent,也就是标记
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	//只有到最后一个递归，uf.parent[x] ＝= x时，才会执行下一句
	return uf.parent[x]
}
func (uf *unionFind) union(x, y int) bool {
	//若两点在同一连通分量中，则返回false,否则进行连接(合并)，并返回true
	px, py := uf.find(x), uf.find(y) //查找二者标记
	if px == py {
		return false
	}
	if uf.rank[px] < uf.rank[py] {
		px, py = py, px
	} //让小的树作为子数，比如一个1,一个n,那么自然是１作为子树
	//这个rank记录的是结点个数而不是高度？
	uf.rank[px] += uf.rank[py]
	uf.parent[py] = px
	return true
}
func findRedundantConnection(edges [][]int) []int {
	var ans []int
	n := len(edges)
	uf := newUnionFind(n)
	for _, v := range edges {
		if !uf.union(v[0], v[1]) {
			ans = v
		}
	}
	return ans
}

// @lc code=end

