/*
 * @lc app=leetcode.cn id=1604 lang=golang
 *
 * [1604] 警告一小时内使用相同员工卡大于等于三次的人
 */

// @lc code=start
package main

import (
	"sort"
	"strconv"
)

// hm1 > hm2
func Lower(hm1, hm2 string) bool {
	h1 := hm1[:2]
	m1 := hm1[3:]
	h2 := hm2[:2]
	m2 := hm2[3:]
	if h1 < h2 {
		return true
	}
	if h1 == h2 && m1 < m2 {
		return true
	}
	return false
}

// hm1 > hm2
func duration(hm1, hm2 string) int {
	h1, _ := strconv.Atoi(hm1[:2])
	m1, _ := strconv.Atoi(hm1[3:])
	h2, _ := strconv.Atoi(hm2[:2])
	m2, _ := strconv.Atoi(hm2[3:])

	if h1 == h2 {
		return m1 - m2
	}
	return (h1-1-h2)*60 + 60 - m2 + m1
}

func alertNames(keyName []string, keyTime []string) []string {
	l := len(keyName)
	if l < 3 {
		return []string{}
	}

	staff := map[string][]string{}
	for i := 0; i < l; i++ {
		name := keyName[i]
		time := keyTime[i]
		staff[name] = append(staff[name], time)
	}

	result := []string{}
	for name, times := range staff {
		if len(times) < 3 {
			continue
		}
		sort.Slice(times, func(i, j int) bool {
			return Lower(times[i], times[j])
		})
		for i := 0; i <= len(times)-3; i++ {
			if duration(times[i+2], times[i]) <= 60 {
				result = append(result, name)
				break
			}
		}
	}

	sort.Strings(result)

	return result
}

// @lc code=end
