package leetcode_go

import "strings"

func intToRoman(num int) string {
	result := make([]string,0);
	units := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	values := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	for i := len(units) - 1; i >= 0 ; i--  {
		for num - values[i] >= 0 {
			result = append(result, units[i])
			num -= values[i]
		}
	}

	return strings.Join(result, "")
}