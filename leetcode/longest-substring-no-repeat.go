package leetcode

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	runes := []rune(s)
	length := 0
	for i, j := 0, 0; i < len(runes) && j < len(runes) && i <= j; {
		if isCharRepeat(m, runes[j]) {
			if length < len(m) {
				length = len(m)
			}
			i = m[runes[j]] + 1
			for k, v := range m {
				if v < i {
					delete(m, k)
				}
			}
		} else {
			m[runes[j]] = j
			j++
		}
	}

	if length < len(m) {
		length = len(m)
	}

	return length
}

func isCharRepeat(m map[rune]int, c rune) bool {
	_, ok := m[c]
	if ok {
		return true
	}

	return false
}
