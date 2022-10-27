package solve0127

/*
*
@index 127
@title 单词接龙
@difficulty 中等
@tags breadth-first-search
@draft false
@link https://leetcode-cn.com/problems/word-ladder/
@frontendId 127.
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !(contains(endWord, wordList)) {
		return 0
	}
	nearMap := convertMap(beginWord, wordList)
	step := 1
	queue := []string{beginWord}
	mm := make(map[string]struct{}, 0)
	mm[beginWord] = struct{}{}
	for len(queue) > 0 {
		next := make([]string, 0)
		for _, w := range queue {
			if w == endWord {
				return step
			}
			for _, idx := range nearMap[w] {
				nextWord := wordList[idx]
				if _, ok := mm[nextWord]; !ok {
					mm[nextWord] = struct{}{}
					next = append(next, nextWord)
				}
			}
		}
		queue = next
		step++
	}
	return 0
}

func LadderLength(beginWord string, endWord string, wordList []string) int {
	return ladderLength(beginWord, endWord, wordList)
}

func contains(endWord string, wordList []string) bool {
	for _, w := range wordList {
		if w == endWord {
			return true
		}
	}
	return false
}

func convertMap(beginWord string, wordList []string) map[string][]int {
	res := make(map[string][]int, len(wordList)+1)
	indexes := make([]int, 0)
	for i, w := range wordList {
		if canConvert(beginWord, w) {
			indexes = append(indexes, i)
		}
	}
	res[beginWord] = indexes

	for i, w := range wordList {
		indexes := make([]int, 0)
		for j, ww := range wordList {
			if i == j {
				continue
			}
			if canConvert(w, ww) {
				indexes = append(indexes, j)
			}
		}
		res[w] = indexes
	}
	return res
}

func canConvert(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	k := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			k++
		}
		if k > 1 {
			return false
		}
	}
	return true
}
