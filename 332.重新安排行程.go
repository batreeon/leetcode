/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 */

// @lc code=start
package main
import "sort"

func findItinerary(tickets [][]string) []string {

	/*
	// 原以为是拓扑排序，但其实第二个示例就不符合
	order := []string{}
	in := map[string]int{}
	next := map[string][]string{}

	for _,ticket := range tickets {
		if _,ok := in[ticket[0]] ; !ok {
			in[ticket[0]] = 0
		}
		in[ticket[1]]++
		next[ticket[0]] = append(next[ticket[0]],ticket[1])
	}

	for k,inV := range in {
		if inV == 0 {
			order = append(order,k)
			delete(in,k)
		}
	}
	sort.Sort(sort.StringSlice(order))
	start := 0
	for {
		res := []string{}
		for i := start ; i < len(order) ; i++ {
			del := order[i]
			nextNodes := next[del]
			for _,node := range nextNodes {
				in[node]--
				if in[node] == 0 {
					res = append(res,node)
				}
				delete(in,node)
			}
		}
		if len(res) == 0 {
			break
		}
		sort.Sort(sort.StringSlice(res))
		order = append(order,res...)
		start += len(res)
	}
	return order
	*/


	// 超级慢的回溯啊啊啊
	next := map[string][]string{}
	for _,ticket := range tickets {
		// if _,ok := next[ticket[0]] ; !ok {
		// 	next[ticket[0]] = []string{}
		// }
		next[ticket[0]] = append(next[ticket[0]],ticket[1])
	}
	for _,v := range next {
		sort.Strings(v)
	}

	/*
	result := []string{"JFK"}
	var backtrack func(last string) bool
	backtrack = func(last string) bool {
		if len(result)-1 == len(tickets) {
			return true
		}
		for i := 0 ; i < len(next[last]) ; i++ {
			nextCity := next[last][i]
			result = append(result,nextCity)
			copy(next[last][i:],next[last][i+1:])
			next[last] = next[last][:len(next[last])-1]
			if backtrack(nextCity) == true {
				return true
			}
			result = result[:len(result)-1]
			next[last] = append(next[last],"")
			copy(next[last][i+1:],next[last][i:])
			next[last][i] = nextCity
		}
		return false
	}

	backtrack("JFK")
	*/

	result := []string{}
	var dfs func(cur string)
	dfs = func(cur string) {
		for {
			if v,ok := next[cur] ; !ok || len(v) == 0 {
				break
			}
			temp := next[cur][0]
			next[cur] = next[cur][1:]
			dfs(temp)
		}
		result = append(result,cur)
	}
	
	dfs("JFK")
	for i,j := 0,len(result)-1 ; i < j ; i,j = i+1,j-1 {
		result[i],result[j] = result[j],result[i]
	}
	return result
}
// @lc code=end

