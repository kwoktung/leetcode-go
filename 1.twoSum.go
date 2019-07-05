package leetcode_go

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		m[val] = i
	}
	for i := 0;  i< len(nums); i++  {
		current := nums[i]
		if v, ok := m[target-current]; ok && v != i {
			return []int{i, v}
		}
	}
	return []int{}
}

func twoSumB(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		if v, ok := m[target-val]; ok && v != i {
			return []int{v, i}
		}
		m[val] = i
	}
	return []int{}
}