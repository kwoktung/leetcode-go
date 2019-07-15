package leetcode_go

func longestCommonPrefix(strs []string) string {
	longest := ""
	if len(strs) == 0 {
		return longest
	} else if len(strs) == 1 {
		return strs[0]
	}
	word := strs[0]
	for i, l := 0, len(word) ; i < l ; i++  {
		char := word[i]
		for _, item := range strs {
			if len(item) <= i || item[i] != char {
				goto back
			}
		}
		longest += string(char)
	}
back:
	return longest
}