package leetcode_go

import "math"

func reverse(x int) int {
	ans := 0;
	for x!= 0 {
		ans = ans * 10 + x%10
		x = x/10
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return 0
	} else {
		return ans
	}
}