package leetcode_go

func searchRange(nums []int, target int) []int {
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0,0}
		}
	}
	start, end := -1, -1
	head, tail := 0, len(nums) - 1
	for head <= tail {
		mid := (head + tail) / 2
		n := nums[mid]
		if n == target {
			start, end = mid, mid
			for start-1 >= 0 && nums[start-1] == target  {
				start -= 1
			}
			for end + 1 < len(nums) && nums[end+1] == target  {
				end += 1
			}
			break
		} else if n < target {
			head = mid + 1
		} else {
			tail = mid - 1
		}
	}
	return []int{start, end}
}