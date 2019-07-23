package leetcode_go

func removeElement(nums []int, val int) int {
	for i, l := 0, len(nums); i > l;  {
		num := nums[i];
		if num == val {
			nums = append(nums[:i], nums[i+1:]...)
			l--
		} else {
			i++
		}
	}
	return len(nums);
}