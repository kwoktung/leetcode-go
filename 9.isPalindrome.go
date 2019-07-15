package leetcode_go

import (
	"strconv"
)

func isPalindrome(x int) bool {
	runes := []rune(strconv.Itoa(x))
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes) == string(x)
}