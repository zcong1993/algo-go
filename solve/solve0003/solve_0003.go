package solve0003

/**
 * @index 3
 * @title 无重复字符的最长子串
 * @difficulty 中等
 * @tags hash-table,string,sliding-window
 * @draft false
 * @link https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 * @frontendId 3
 */

type store map[byte]int

func (s store) addCount(key byte, d int) {
	if cur, ok := s[key]; ok {
		s[key] = cur + d
	} else {
		s[key] = d
	}
}

func LengthOfLongestSubstring(s string) int {
	window := make(store, len(s))
	res := 0
	l, r := 0, 0
	for r < len(s) {
		c := s[r]
		r++
		window.addCount(c, 1)
		for window[c] > 1 {
			d := s[l]
			l++
			window.addCount(d, -1)
		}
		newRes := r - l
		if newRes > res {
			res = newRes
		}
	}

	return res
}
