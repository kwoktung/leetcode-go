package leetcode_go

func romanToInt(s string) int {
	sum, prev := 0, 10000;
	alphabet := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for _, char := range s {
		value := alphabet[byte(char)]
		if value > prev {
			sum += value - 2 * prev
		} else {
			sum += value
		}
		prev = value
	}
	return sum
}