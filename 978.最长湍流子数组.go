/*
 * @lc app=leetcode.cn id=978 lang=golang
 *
 * [978] 最长湍流子数组
 */

// @lc code=start
func maxTurbulenceSize(arr []int) int {
	if len(arr) == 1 {
		return 1
	}
	left,right := 0,1
	up := 0 //标记变化趋势，增大置1，减小置-1，不变置0
	maxLen := 1

	max := func(a,b int) int {
		if a > b {
			return a
		}
		return b
	}
	
	for ; right < len(arr) ; right++ {
		if arr[right] > arr[right-1] {
			if up == 1 {
				left = right-1
			}else{
				up = 1
				maxLen = max(maxLen,right-left+1)
			}
		}else if arr[right] < arr[right-1] {
			if up == -1{
				left = right-1
			}else{
				up = -1
				maxLen = max(maxLen,right-left+1)
			}
		}else if arr[right] == arr[right-1] {
			if up != 0 {
				up = 0
			}
				left = right
		}

		//上面的和下面注释掉的完全一样，只不过改了下顺序
		/*
		if up == 0 {
			if arr[right] > arr[right-1] {
				up = 1
				maxLen = max(maxLen,right-left+1)
			}else if arr[right] < arr[right-1] {
				up = -1
				maxLen = max(maxLen,right-left+1)
			}else if arr[right] == arr[right-1] {
				left = right
			}
		}else if up == -1 {
			if arr[right] > arr[right-1] {
				up = 1
				maxLen = max(maxLen,right-left+1)
			}else if arr[right] < arr[right-1] {
				left = right-1
			}else if arr[right] == arr[right-1] {
				up = 0
				left = right
			}
		}else if up == 1 {
			if arr[right] > arr[right-1] {
				left = right-1
			}else if arr[right] < arr[right-1] {
				up = -1
				maxLen = max(maxLen,right-left+1)
			}else if arr[right] == arr[right-1] {
				up = 0
				left = right
			}
		}
		*/
	}
	return maxLen
}
// @lc code=end

