package leetcode_go

func lengthOfLongestSubstring(s string) int {
	kv := make(map[byte] int)
	max := 0
	for start, end := 0, 0; end < len(s); end++  {
		key := s[end]
		if v, ok := kv[key]; ok && v >= start {
			start = kv[key] + 1
		}
		if end - start + 1 > max {
			max = end - start + 1
		}
		kv[key] = end
	}
	return max
}
