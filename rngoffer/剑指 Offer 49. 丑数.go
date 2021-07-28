package main
func nthUglyNumber(n int) int {
	res := []int{1}
	i2,i3,i5 := 0,0,0
	min := func(x,y,z int) int {
		if x > y {
			x,y = y,x
		}
		if x > z {
			return z
		}
		return x
	}
	for {
		if len(res) == n {
			return res[n-1]
		}
		next := min(2*res[i2],3*res[i3],5*res[i5])
		res = append(res,next)
		for 2*res[i2] <= next {
			i2++
		}
		for 3*res[i3] <= next {
			i3++
		}
		for 5*res[i5] <= next {
			i5++
		}
	}
}