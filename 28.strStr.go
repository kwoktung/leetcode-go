package leetcode_go


func strStr(haystack string, needle string) int {
	if len(needle) == 0 { return 0 }
	l, result := len(needle), 0

	for len(haystack) >= l {
		word := haystack[0:l];
		if word == needle {
			return result
		} else {
			result += 1
			haystack = haystack[1:]
		}
	}
	return -1
}