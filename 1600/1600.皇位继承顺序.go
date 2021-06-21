/*
 * @lc app=leetcode.cn id=1600 lang=golang
 *
 * [1600] 皇位继承顺序
 */

// @lc code=start
package main
// import "fmt"


type ThroneInheritance struct {
	king string
	children map[string][]string
	parent map[string]string
	death map[string]bool
	total int
}


func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		kingName,
		make(map[string][]string),
		make(map[string]string),
		make(map[string]bool),
		1,
	}
}


func (this *ThroneInheritance) Birth(parentName string, childName string)  {
	if _,ok := this.children[parentName] ; !ok {
		this.children[parentName] = []string{}
	}
	this.children[parentName] = append(this.children[parentName],childName)
	this.parent[childName] = parentName
	this.total++
}


func (this *ThroneInheritance) Death(name string)  {
	this.death[name] = true
}


func (this *ThroneInheritance) GetInheritanceOrder() []string {
	var successor func(x string,curorder *[]string,curorderm *map[string]bool)
	successor = func(x string,curorder *[]string,curorderm *map[string]bool) {
		if len(*curorder) == this.total {
			return
		}
		pass := false
		if children,ok := this.children[x] ; ok {
			for _,child := range children {
				if !(*curorderm)[child] {
					pass = true
					break
				}
			}
		}
		if !pass {
			if x == this.king {
				return
			}else{
				successor(this.parent[x],curorder,curorderm)
			}
		}else{
			for _,child := range this.children[x] {
				if !(*curorderm)[child] {
					*curorder = append((*curorder),child)
					(*curorderm)[child] = true
					successor(child,curorder,curorderm)
				}
			}
		}
		return
	}

	curOrder := []string{this.king}
	successor(this.king,&curOrder,&(map[string]bool{this.king:true}))
	result := []string{}
	for i := range curOrder {
		if !this.death[curOrder[i]] {
			result = append(result,curOrder[i])
		}
	}
	return result
}

// func main() {
// 	t := Constructor("king")
// 	fmt.Println(t.GetInheritanceOrder())
// 	t.Birth("king","clyde")
// 	t.Birth("clyde","shannon")
// 	t.Birth("shannon","scott")
// 	t.Birth("king","keith")
// 	fmt.Println(t.GetInheritanceOrder())
// 	t.Birth("clyde","joseph")
// 	fmt.Println(t.GetInheritanceOrder())
// }

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
// @lc code=end

