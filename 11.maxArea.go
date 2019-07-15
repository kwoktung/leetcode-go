package leetcode_go

func min (a,b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	max := 0;
	head, tail := 0, len(height) - 1
	for head < tail  {
		headH, tailH := height[head], height[tail];
		area :=  (tail - head) * min(height[head], height[tail]);
		if area > max {
			max = area
		}
		if headH < tailH {
			head++
		} else {
			tail--
		}
	}
	return max
}