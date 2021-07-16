func verifyPostorder(postorder []int) bool {
    var verify func(i,j int) bool
	verify = func(i,j int) bool {
		// 为什么有大于号，考虑只有右子树没有左子树的情况，m-1 < i，那么他就会无限的递归下去
		if i >= j {
			return true
		}
		p := i
		for postorder[p] < postorder[j] {p++}
		m := p
		for ; p < j ; p++ {
			if postorder[p] < postorder[j] {
				return false
			}
		}
		return verify(i,m-1) && verify(m,j-1)
	}
    return verify(0,len(postorder)-1)
}