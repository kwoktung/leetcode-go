package leetcode_go


func searchInsert(nums []int, target int) int {
	for i, l := 0, len(nums); i < l ; i++  {
		c := nums[i];
		if c >= target {
			return i
		}
	}
	return len(nums)
}