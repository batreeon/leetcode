/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 */

// @lc code=start
// package main 
// import (
// 	"fmt"
// )

// func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
// 	m := map[string]float64{}
// 	letterM := map[string]bool{}
// 	for i,v := range equations {
// 		m[v[0]+"/"+v[1]] = values[i]
// 		letterM[v[0]] = true
// 		letterM[v[1]] = true
// 	}
// 	ans := []float64{}
// 	for _,v := range queries {
// 		ans1 := calcEquationRecur(equations,m,letterM,v)
// 		if ans1 != -1{
// 			ans = append(ans,ans1)
// 		}else{
// 			ans2 := calcEquationRecur(equations,m,letterM,[]string{v[1],v[0]})
// 			ans = append(ans,1/ans2)
// 			//这里没有判断ans2是否为-1,因为不是-1，则应添加1/ans2，
// 			//若是-1，则应添加-1,刚好1/-1等于-1，判不判断都一样
// 		}
// 	}
// 	return ans
// }
// func calcEquationRecur(equations [][]string, m map[string]float64, letterM map[string]bool, f []string) float64 {
	
// 	if !letterM[f[0]] || !letterM[f[1]] {
// 		return -1
// 	}else if f[0] == f[1] {
// 		return 1
// 	}else if d,ok := m[f[0]+"/"+f[1]]; ok {
// 		return d
// 	}else if d,ok := m[f[1]+"/"+f[0]]; ok {
// 		return 1/d
// 	}else{
// 		for _,f1 := range equations {
// 			if f1[0] == f[0] {
// 				f2 := []string{f1[1],f[1]}
// 				f2Value := calcEquationRecur(equations,m,letterM,f2)
// 				_f2 := []string{f[1],f1[1]}
// 				_f2Value := calcEquationRecur(equations,m,letterM,_f2)
// 				if f2Value == -1 && _f2Value == -1 {
// 					break
// 				}else if f2Value != -1{
// 					return m[f1[0]+"/"+f1[1]]*f2Value
// 				}else{
// 					return m[f1[0]+"/"+f1[1]]*(1/_f2Value)
// 				}
// 			}
// 		}
// 		for _,f1 := range equations {
// 			if f1[1] == f[1] {
// 				f2 := []string{f[0],f1[0]}
// 				f2Value := calcEquationRecur(equations,m,letterM,f2)
// 				_f2 := []string{f1[0],f[0]}
// 				_f2Value := calcEquationRecur(equations,m,letterM,_f2)
// 				if f2Value == -1 && _f2Value == -1 {
// 					break
// 				}else if f2Value != -1{
// 					return f2Value*m[f1[0]+"/"+f1[1]]
// 				}else{
// 					return (1/_f2Value)*m[f1[0]+"/"+f1[1]]
// 				}
// 			}
// 		}
// 	}
// 	return -1
// }
// func main() {
// 	e := {["x1","x2"],["x2","x3"],["x1","x4"],["x2","x5"]}
// 	v := {3.0,0.5,3.4,5.6}
// 	q := {["x2","x4"],["x1","x5"],["x1","x3"],["x5","x5"],["x5","x1"],["x3","x4"],["x4","x3"],["x6","x6"],["x0","x0"]}
// 	calcEquation(e,v,q)
// }

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	distance := map[string]float64{}//记录边
	nearV := map[string][]string{}//邻接矩阵
	for i,v := range equations {
		distance[v[0]+"/"+v[1]] = values[i]
		distance[v[1]+"/"+v[0]] = 1/values[i]
		nearV[v[0]] = append(nearV[v[0]],v[1])
		nearV[v[1]] = append(nearV[v[1]],v[0])
	}

	visited := map[string]bool{}
	var calcEquationDfs func(f []string, visited map[string]bool) float64
	calcEquationDfs = func (f []string, visited map[string]bool) float64 {
		if _,ok := nearV[f[0]]; !ok {
			return 0
		}
		if _,ok := nearV[f[1]]; !ok {
			return 0
		}//有节点不存在
		if visited[f[0]] {
			return 0
		}//已经访问过了，避免产生环路
		if f[0] == f[1] {
			return 1
		}//分子分母相同
		if d,ok := distance[f[0]+"/"+f[1]]; ok {
			return d
		}
		// if d,ok := distance[f[1]+"/"+f[0]]; ok {
		// 	return 1/d
		// }//不能用这个来替代在distance中添加反向边，
		//因为前面构造图默认邻接矩阵记录的是无向图
		visited[f[0]] = true
		res := 0.0
		for _,tmp := range nearV[f[0]] {
			res = calcEquationDfs([]string{tmp,f[1]},visited) * distance[f[0]+"/"+tmp]
			if res != 0 {
				distance[f[0]+"/"+f[1]] = res
				break
			}
		}
		delete(visited,f[0])
		return res
	}
	
	ans := []float64{}
	for _,v := range queries {
		tmp := calcEquationDfs(v,visited)
		if tmp == 0 {
			tmp = -1
		}
		ans = append(ans,tmp)
	}
	return ans
}

// func calcEquationRecur(equations [][]string, m map[string]float64, letterM map[string]bool, f []string) float64 {
	
// 	if !letterM[f[0]] || !letterM[f[1]] {
// 		return -1
// 	}else if f[0] == f[1] {
// 		return 1
// 	}else if d,ok := m[f[0]+"/"+f[1]]; ok {
// 		return d
// 	}else if d,ok := m[f[1]+"/"+f[0]]; ok {
// 		return 1/d
// 	}else{
// 		for _,f1 := range equations {
// 			if f1[0] == f[0] {
// 				f2 := []string{f1[1],f[1]}
// 				f2Value := calcEquationRecur(equations,m,letterM,f2)
// 				_f2 := []string{f[1],f1[1]}
// 				_f2Value := calcEquationRecur(equations,m,letterM,_f2)
// 				if f2Value == -1 && _f2Value == -1 {
// 					break
// 				}else if f2Value != -1{
// 					return m[f1[0]+"/"+f1[1]]*f2Value
// 				}else{
// 					return m[f1[0]+"/"+f1[1]]*(1/_f2Value)
// 				}
// 			}
// 		}
// 		for _,f1 := range equations {
// 			if f1[1] == f[1] {
// 				f2 := []string{f[0],f1[0]}
// 				f2Value := calcEquationRecur(equations,m,letterM,f2)
// 				_f2 := []string{f1[0],f[0]}
// 				_f2Value := calcEquationRecur(equations,m,letterM,_f2)
// 				if f2Value == -1 && _f2Value == -1 {
// 					break
// 				}else if f2Value != -1{
// 					return f2Value*m[f1[0]+"/"+f1[1]]
// 				}else{
// 					return (1/_f2Value)*m[f1[0]+"/"+f1[1]]
// 				}
// 			}
// 		}
// 	}
// 	return -1
// }

// @lc code=end

