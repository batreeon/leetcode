func fib(n int) int {
	if n <= 1 {
		return n
	}
	f0,f1 := 0,1
	for n - 1 > 0 {
		f0,f1 = f1,f0+f1
		n--
	}
	return f1
}