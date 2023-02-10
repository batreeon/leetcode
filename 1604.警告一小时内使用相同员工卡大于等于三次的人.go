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

type hm string

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

func duration(hm1, hm2 string) int {
	h1,_ := strconv.Atoi(hm1[:2])
	m1,_ := strconv.Atoi(hm1[3:])
	h2,_ := strconv.Atoi(hm2[:2])
	m2,_ := strconv.Atoi(hm2[3:])

	if h1 == h2 {
		return m1 - m2
	}
	return (h1 - 1 - h2) * 60 + 60 - m2 + m1 
}

func alertNames(keyName []string, keyTime []string) []string {
	staff := map[string][]string{}

	l := len(keyName)
	result := []string{}
	if l < 3 {
		return result
	}

	resultM := map[string]struct{}{}
	for i := 0; i < l; i++ {
		name := keyName[i]
		time := keyTime[i]

		if _, ok := staff[name]; !ok {
			staff[name] = []string{time}
		}else{
			if Lower(time, staff[name][len(staff[name])-1]) {
				staff[name] = []string{time}
				continue
			}
			staff[name] = append(staff[name], time)
			index := sort.Search(len(staff[name]), func(i int) bool {
				return duration(time, staff[name][i]) <= 60
			})
			staff[name] = staff[name][index:]
			if len(staff[name]) >= 3 {
				resultM[name] = struct{}{}
			}
		}
	}

	for k := range resultM {
		result = append(result, k)
	}

	sort.Strings(result)

	return result
}

// @lc code=end
