package leetcode

import "math"

// https://leetcode.cn/contest/weekly-contest-408/problems/count-the-number-of-substrings-with-dominant-ones/

func numberOfSubstrings(s string) int {
	next0 := make(map[int]int, len(s))
	var ti int = len(s)
	for i := len(s) - 1; i >= 0; i-- {
		next0[i] = ti
		if s[i] == '0' {
			ti = i
		}
	}
	next0[len(s)] = -1
	maxC0 := int(math.Sqrt(float64(len(s))))
	maxC0 = min(maxC0+5, len(s))
	var right int
	var ans int = 0

	// 枚举子串的左端点 O(4*10^4)
	for left := 0; left < len(s); left++ {
		right = next0[left]
		c0 := 0
		if s[left] == '0' {
			c0 = 1
		}
		var lastRight int
		// 枚举字串的 0 数量 O(2*10^2)
		for c0 < maxC0 && right != -1 {
			c1 := right - left - c0
			// 如果字符串中 1 的数量 大于或等于 0 的数量的 平方，则认为该字符串是一个 1 显著 的字符串
			if c1 >= c0*c0 {
				a1 := c1 - c0*c0
				if c0 > 0 {
					a1++
				}
				ans += min(a1, right-lastRight)
			}
			c0++
			lastRight = right
			right = next0[right]
		}
	}
	return ans
}
