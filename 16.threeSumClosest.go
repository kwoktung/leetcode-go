package leetcode_go

import "sort"

func abs (n int) int {
	if n < 0 { return -n }
	return n
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sort.Ints(nums);
	result := nums[0] + nums[1] + nums[len(nums) -1]
	last := abs(nums[0] + nums[1] + nums[len(nums) -1] - target);
	for i, l := 0, len(nums); i < l - 1; i++  {
		left, right := i + 1, len(nums) - 1;
		for left < right {
			current := abs(nums[i] + nums[left] + nums[right] - target);
			if current == 0 {
				return nums[i] + nums[left] + nums[right]
			} else if current < last {
				last = current
				result = nums[i] + nums[left] + nums[right]
			}
			if nums[i] + nums[left] + nums[right] - target > 0 {
				right -= 1;
			} else {
				left += 1;
			}
		}
	}
	return result
}