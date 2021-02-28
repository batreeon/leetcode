/*
 * @lc app=leetcode.cn id=896 lang=golang
 *
 * [896] 单调数列
 */

// @lc code=start
func isMonotonic(A []int) bool {
	/*通过了，速度不错但占内存
	n := len(A)
	if n < 3 {
		return true
	}
	up,down := false,false
	i,j := 0,1
	for ; j < n && ( (!up && !down) || (up && A[i] <= A[j]) || (down && A[i] >= A[j]) ) ; i,j = j,j+1 {
		if A[i] != A[j] {
			if A[i] < A[j] {
				up = true
			}else{
				down = true
			}
		}
	} 
	if j == n {
		return true
	}
	return false
	*/

	/*
	依然耗内存
	n := len(A)
	if n < 3 {
		return true
	}
	j := 1
	for ; A[j-1] == A[j] ; {
		if j == n-1 {
			return true
		}
		j++
	}
	up := A[j-1] < A[j]
	for ; up == (A[j-1] < A[j]) || A[j-1] == A[j] ; {
		if j == n-1 {
			return true
		}
		j++
	}
	return false
	*/

	/*
	耗时太大
	a := sort.IntSlice(A)
	return sort.IsSorted(a) || sort.IsSorted(sort.Reverse(a))
	*/

	/*
	和前面两种差别不大
	n := len(A)
	inc,dec := true,true
	j := 1
	for inc || dec {
		if j >= n {
			break
		}
		if A[j-1] < A[j] {
			dec = false
		}
		if A[j-1] > A[j] {
			inc = false
		}
		j++
	} 
	if !inc && !dec {
		return false
	}
	return true
	*/
}
// @lc code=end