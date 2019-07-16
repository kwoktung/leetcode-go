package leetcode_go

func isValid(s string) bool {
	m := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	list := make([]byte, 0)
	for i, l := 0, len(s); i < l; i++ {
		c := s[i]
		if v, ok := m[c]; ok && len(list) > 0 {
			current := list[len(list)-1]
			list = list[:len(list)-1]
			if current != v {
				return false
			}
		} else {
			list = append(list, c)
		}
	}
	return len(list) == 0
}