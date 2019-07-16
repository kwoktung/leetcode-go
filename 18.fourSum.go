package leetcode_go

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var answer [][]int
	sort.Ints(nums)
	for i := 0 ; i < len(nums) - 3; i++ {
		if i > 0 && nums[i] == nums[i - 1] { continue }
		if nums[i] +nums[i+1] + nums[i+2] + nums[i+3] > target { break }
		for j := i + 1; j < len(nums) - 2; j++ {
			if j - i > 1 && nums[j] == nums[j-1] { continue }
			if nums[i] + nums[j] + nums[j+1] + nums[j+2] > target { break }
			if nums[i] + nums[j] + nums[len(nums)- 2] + nums[len(nums) - 1] < target { continue }
			head, tail := j + 1, len(nums) - 1
			for head < tail {
				n := nums[i] + nums[j] + nums[head] + nums[tail]
				if n == target {
					answer = append(answer, []int{nums[i], nums[j], nums[head],nums[tail] })
					for head < tail && nums[head] == nums[head+1]  {
						head++
					}
					for head < tail && nums[tail] == nums[tail-1] {
						tail--
					}
					head++
					tail--
				} else {
					if n < target {
						head++
					} else {
						tail--
					}
				}
			}
		}
	}
	return answer
}
