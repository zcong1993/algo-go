package solve0076

import "math"

/**
 * @index 76
 * @title 最小覆盖子串
 * @difficulty 困难
 * @tags hash-table,string,sliding-window
 * @draft false
 * @link https://leetcode-cn.com/problems/minimum-window-substring/
 * @frontendId 76
 */

type mp map[byte]int

func (m mp) addCount(key byte, d int) {
	if cur, ok := m[key]; ok {
		m[key] = cur + d
	} else {
		m[key] = d
	}
}

func (m mp) has(key byte) bool {
	_, ok := m[key]
	return ok
}

func MinWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	window := make(mp, len(t))
	need := make(mp, len(t))
	for i := range t {
		need.addCount(t[i], 1)
	}

	var start int
	ll := math.MaxInt32
	needLen := len(need)
	valid := 0
	l, r := 0, 0
	for r < len(s) {
		c := s[r]
		r++
		if need.has(c) {
			window.addCount(c, 1)
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == needLen {
			newRes := r - l
			if newRes < ll {
				start = l
				ll = newRes
			}
			d := s[l]
			l++
			if need.has(d) {
				window.addCount(d, -1)
				if window[d] < need[d] {
					valid--
				}
			}
		}
	}

	if ll == math.MaxInt32 {
		return ""
	}

	return s[start : start+ll]
}
