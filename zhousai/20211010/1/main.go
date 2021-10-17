// 给你三个整数数组 nums1、nums2 和 nums3 ，请你构造并返回一个 不同 数组，且由 至少 在 两个 数组中出现的所有值组成。数组中的元素可以按 任意 顺序排列

package main

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m := map[int]int{}
	for _,num1 := range nums1 {
		if _,ok := m[num1]; !ok {
			m[num1] = 1
		}
	}

	for _,num2 := range nums2 {
		if n,ok := m[num2]; !ok || n == 1 {
			m[num2] += 2
		}
	}

	for _,num3 := range nums3 {
		if n,ok := m[num3]; !ok || n == 1 || n == 2 {
			m[num3] += 4
		}
	}

	result := []int{}
	for k,v := range m {
		if v == 3 || v == 6 || v == 5 {
			result = append(result, k)
		}
	}

	return result
}