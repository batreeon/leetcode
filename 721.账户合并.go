/*
 * @lc app=leetcode.cn id=721 lang=golang
 *
 * [721] 账户合并
 */

// @lc code=start
// func accountsMerge(accounts [][]string) [][]string {
// 	parent := map[string]int{}
// 	account := map[int][]string{}
// 	for i , v := range accounts {
// 		nowParent := i
// 		for j := 1 ; j < len(v) ; j++ {
// 			if p,ok := parent[v[j]];ok{
// 				if p < nowParent {
// 					nowParent = p
// 				}
// 			} 
// 		}
// 		for j := 1 ; j < len(v) ; j++ {
// 			if p,ok := parent[v[j]];!ok{
// 				account[nowParent] = append(account[nowParent],v[j])
// 				sort.Strings(account[nowParent])
// 				parent[v[j]] = nowParent
// 			}else{
// 				if p != nowParent {
// 					for _,vv := range account[p] {
// 						account[nowParent] = append(account[nowParent],vv)
// 						sort.Strings(account[nowParent])
// 					}
// 					delete(account,p)
// 					parent[v[j]] = nowParent
// 				}
// 			}
// 		}
// 	}
// 	ans := [][]string{}
// 	for k,v := range account {
// 		oneaccount := []string{}
// 		oneaccount = append(oneaccount,accounts[k][0])
// 		for _,vv := range v {
// 			oneaccount = append(oneaccount,vv)
// 		}
// 		ans = append(ans,oneaccount)
// 	}
// 	return ans
// }
func accountsMerge(accounts [][]string) (ans [][]string) {
    emailToIndex := map[string]int{}
    emailToName := map[string]string{}
    for _, account := range accounts {
        name := account[0]
        for _, email := range account[1:] {
            if _, has := emailToIndex[email]; !has {
                emailToIndex[email] = len(emailToIndex) //记录所有不同的email地址
                emailToName[email] = name   //记录email对应的用户名
            }
        }
    }

    parent := make([]int, len(emailToIndex))
    for i := range parent {
        parent[i] = i
    }
    var find func(int) int  //函数值类型
    find = func(x int) int {
        if parent[x] != x {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }
    union := func(from, to int) {
        parent[find(from)] = find(to)   //没有考虑优化树的深度
    }

    for _, account := range accounts {
        firstIndex := emailToIndex[account[1]]
        for _, email := range account[2:] {
            union(emailToIndex[email], firstIndex)
        }
    }//属于同一账户的邮箱就拥有同样的int值

    indexToEmails := map[int][]string{}
    for email, index := range emailToIndex {
        index = find(index)//找的firstIndex,也就是最父亲的那个结点
        indexToEmails[index] = append(indexToEmails[index], email)
    }

    for _, emails := range indexToEmails {
        sort.Strings(emails)
        //emails[0]表示mails序列中排第一的那个，emialToNmae记录了每个email对应的账户名
        account := append([]string{emailToName[emails[0]]}, emails...)
        ans = append(ans, account)
    }
    return
}
// @lc code=end

