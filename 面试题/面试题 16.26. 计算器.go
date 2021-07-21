package main

import "strings"

func calculate(s string) int {
	m := map[byte]int{'+':1,'-':1,'*':2,'/':2}
	s = strings.ReplaceAll(s," ","")
	ops,nums := []byte{},[]int{}
	cal := func() {
		if len(nums) < 2 || len(ops) == 0 {
			return
		}
		b,a := nums[len(nums)-1],nums[len(nums)-2]
		nums = nums[:len(nums)-2]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		res := 0
		if op == '+' {res = a+b}else 
		if op == '-' {res = a-b}else
		if op == '*' {res = a*b}else
		if op == '/' {res = a/b}
		nums = append(nums,res)
	}
	for i := 0 ; i < len(s) ; {
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			nums = append(nums,num)
		}else{
			for len(ops) > 0 {
				prev := ops[len(ops)-1]
				if m[prev] >= m[s[i]] {
					cal()
				}else{
					break
				}
			}
			ops = append(ops,s[i])
			i++
		}
	}
	for len(ops) > 0 {
		cal()
	}
	return nums[0]
}