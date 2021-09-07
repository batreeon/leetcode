package main

func countEval(s string, result int) int {
	n := len(s)

	dp0 := make([][]int,n/2+1)
	dp1 := make([][]int,n/2+1)
	for i := 0 ; i < n/2+1 ; i++ {
		dp0[i] = make([]int,n/2+1)
		if s[i*2] == '0' {
			dp0[i][i] = 1
		}
		dp1[i] = make([]int,n/2+1)
		if s[i*2] == '1' {
			dp1[i][i] = 1
		}
	}
	for l := 2 ; l <= n/2+1 ; l++ {
		for i := 0 ; i < n/2+1-l+1 ; i++ {
			for j := i ; j < i+l-1 ; j++ {
				if s[j*2+1] == '|' {
					dp0[i][i+l-1] += dp0[i][j] * dp0[j+1][i+l-1]
					dp1[i][i+l-1] += dp1[i][j]*(dp1[j+1][i+l-1]+dp0[j+1][i+l-1]) + dp0[i][j]*dp1[j+1][i+l-1]
				}else if s[j*2+1] == '&' {
					dp0[i][i+l-1] += dp0[i][j]*(dp1[j+1][i+l-1]+dp0[j+1][i+l-1]) + dp1[i][j]*dp0[j+1][i+l-1]
					dp1[i][i+l-1] += dp1[i][j] * dp1[j+1][i+l-1]
				}else if s[j*2+1] == '^' {
					dp0[i][i+l-1] += dp0[i][j] * dp0[j+1][i+l-1] + dp1[i][j] * dp1[j+1][i+l-1]
					dp1[i][i+l-1] += dp0[i][j] * dp1[j+1][i+l-1] + dp1[i][j] * dp0[j+1][i+l-1]
				}
			}
		}
	}
	if result == 0 {
		return dp0[0][n/2]
	}
	return dp1[0][n/2]
}
/*
func countEval(s string, result int) int {
	n := len(s)

	dp0 := make([][]int,n/2+1)
	dp1 := make([][]int,n/2+1)
	for i := 0 ; i < n/2+1 ; i++ {
		dp0[i] = make([]int,n/2+1)
		if s[i*2] == '0' {
			dp0[i][i] = 1
		}
		dp1[i] = make([]int,n/2+1)
		if s[i*2] == '1' {
			dp1[i][i] = 1
		}
	}
	for l := 2 ; l <= n/2+1 ; l++ {
		for i := 0 ; i < n/2+1-l+1 ; i++ {
			if l == 2 {
				if s[i*2+1] == '|' {
					if s[i*2] == '0' && s[i*2+2] == '0' {
						dp0[i][i+1] = 1
					}else{
						dp0[i][i+1] = 0
					}
					if s[i*2] == '1' || s[i*2+2] == '1' {
						dp1[i][i+1] = 1
					}else{
						dp1[i][i+1] = 0
					}
				}else if s[i*2+1] == '&' {
					if s[i*2] == '0' || s[i*2+2] == '0' {
						dp0[i][i+1] = 1
					}else{
						dp0[i][i+1] = 0
					}
					if s[i*2] == '1' && s[i*2+2] == '1' {
						dp1[i][i+1] = 1
					}else{
						dp1[i][i+1] = 0
					}
				}else if s[i*2+1] == '^' {
					if s[i*2] == s[i*2+2] {
						dp0[i][i+1] = 1
						dp1[i][i+1] = 0
					}else{
						dp0[i][i+1] = 0
						dp1[i][i+1] = 1
					}
				}
			}else{
				for j := i ; j < i+l-1 ; j++ {
					if s[j*2+1] == '|' {
						dp0[i][i+l-1] += dp0[i][j] * dp0[j+1][i+l-1]
						dp1[i][i+l-1] += dp1[i][j]*(dp1[j+1][i+l-1]+dp0[j+1][i+l-1]) + dp0[i][j]*dp1[j+1][i+l-1]
					}else if s[j*2+1] == '&' {
						dp0[i][i+l-1] += dp0[i][j]*(dp1[j+1][i+l-1]+dp0[j+1][i+l-1]) + dp1[i][j]*dp0[j+1][i+l-1]
						dp1[i][i+l-1] += dp1[i][j] * dp1[j+1][i+l-1]
					}else if s[j*2+1] == '^' {
						dp0[i][i+l-1] += dp0[i][j] * dp0[j+1][i+l-1] + dp1[i][j] * dp1[j+1][i+l-1]
						dp1[i][i+l-1] += dp0[i][j] * dp1[j+1][i+l-1] + dp1[i][j] * dp0[j+1][i+l-1]
					}
				}
			}
		}
	}
	if result == 0 {
		return dp0[0][n/2]
	}
	return dp1[0][n/2]
}
*/