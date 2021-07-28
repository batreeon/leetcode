package main

func cuttingRope(n int) int {
	mod := int(1e9+7)
	if n <= 3 {
		return n-1
	}
	result := 1
	for n > 4 {
		result = result*3%mod
		n -= 3
	}
	return result*n%mod
}