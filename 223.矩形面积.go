/*
 * @lc app=leetcode.cn id=223 lang=golang
 *
 * [223] 矩形面积
 */

// @lc code=start
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	if C <= E || D <= F || A >= G || B >= H {
		return (C-A) * (D-B) + (G-E) * (H-F)
	}
	if A <= E && B <= F && C >= G && D >= H {
		return (C-A) * (D-B)
	}
	if E <= A && F <= B && G >= C && H >= D {
		return (G-E) * (H-F)
	}
	if G <= C && E >= A {
		result := (C-A) * (D-B)
		if F < B {
			result += (G-E) * (B-F)
		}
		if H > D {
			result += (G-E) * (H-D)
		}
		return result
	}
	if C <= G && A >= E {
		result := (G-E) * (H-F)
		if B < F {
			result += (C-A) * (F-B)
		}
		if D > H {
			result += (C-A) * (D-H)
		}
		return result
	}
	if D >= H && B <= F {
		result := (C-A) * (D-B)
		if G > C {
			result += (H-F) * (G-C)
		}
		if E < A {
			result += (H-F) * (A-E)
		}
		return result
	}
	if H >= D && F <= B {
		result := (G-E) * (H-F)
		if C > G {
			result += (D-B) * (C-G)
		}
		if A < E {
			result += (D-B) * (E-A)
		}
		return result
	}

	if C <= G && B >= F {
		return (C-A) * (D-B) + (G-E) * (H-F) - (C-E)*(H-B)
	}
	if C <= G && F >= B {
		return (C-A) * (D-B) + (G-E) * (H-F) - (C-E)*(D-F)
	}
	if A >= E && B >= F {
		return (C-A) * (D-B) + (G-E) * (H-F) - (G-A)*(H-B)
	}
	if A >= E && F >= B {
		return (C-A) * (D-B) + (G-E) * (H-F) - (G-A)*(D-F)
	}
	return -1
}
// @lc code=end

