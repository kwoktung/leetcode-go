package leetcode_go

import (
	"sort"
	)

func threeSum(nums []int) [][]int {
	result := make([][]int, 0);
	sort.Ints(nums)
	for i, len := 0, len(nums); i < len - 1 ; i++  {
		if i > 0 && nums[i] == nums[i-1] { continue }
		target, left, right := i, i + 1, len -1;
		for left < right  {
			targetNum, leftNum, rightNum := nums[target], nums[left], nums[right]
			sum := targetNum + leftNum + rightNum
			if sum == 0 {
				answer := []int{targetNum, leftNum, rightNum}
				result = append(result, answer)
				left++
				right--;
				for nextLeftNum := nums[left]; nextLeftNum == leftNum && left < right;  {
					left++;
					nextLeftNum = nums[left];
				}
				for nextRightNum := nums[right]; nextRightNum == rightNum && left < right;  {
					right--;
					nextRightNum = nums[right];
				}
			} else if sum > 0 {
				right--;
			} else {
				left++;
			}
		}
	}
	return result
}

