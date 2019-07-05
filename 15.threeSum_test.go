package leetcode_go

import "testing"

func equalSlice(a []int, b []int)  bool {
	if a == nil || b== nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true

}

func equal(a[][]int, b [][]int) bool {
	if a == nil || b== nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if(!equalSlice(v, b[i])) {
			return false
		}
	}
	return true
}

func TestThreeSum(t *testing.T) {
	result := threeSum([]int{-1,0,1,2,-1,-4})

	if (!equal(result, [][]int{
		[]int{-1, -1, 2},
		[]int{-1, 0, 1},
	})) {
		t.Error("threeSum []int{-1,0,1,2,-1,-4} Error")
	} else {
		t.Log("threeSum []int{-1,0,1,2,-1,-4} Success")
	}
}