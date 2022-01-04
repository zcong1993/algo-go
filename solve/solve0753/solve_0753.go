package solve0753

/**
 * @index 753
 * @title 打开转盘锁
 * @difficulty 中等
 * @tags breadth-first-search,array,hash-table,string
 * @draft false
 * @link https://leetcode-cn.com/problems/open-the-lock/
 * @frontendId 752
 */

func plus(cur []byte, i int) []byte {
	cp := make([]byte, len(cur))
	copy(cp, cur)
	if cp[i] == '9' {
		cp[i] = '0'
	} else {
		cp[i] += 1
	}
	return cp
}

func minus(cur []byte, i int) []byte {
	cp := make([]byte, len(cur))
	copy(cp, cur)
	if cp[i] == '0' {
		cp[i] = '9'
	} else {
		cp[i] -= 1
	}
	return cp
}

func OpenLock(deadends []string, target string) int {
	res := 0
	visited := make(map[string]struct{})
	queue := [][]byte{[]byte("0000")}
	visited["0000"] = struct{}{}

	deadendsMap := make(map[string]struct{}, len(deadends))
	for _, d := range deadends {
		deadendsMap[d] = struct{}{}
	}

	for len(queue) > 0 {
		var next [][]byte
		for _, cur := range queue {
			cs := string(cur)
			if _, ok := deadendsMap[cs]; ok {
				continue
			}
			if cs == target {
				return res
			}

			for i := 0; i < 4; i++ {
				pn := plus(cur, i)
				pns := string(pn)
				if _, ok := visited[pns]; !ok {
					next = append(next, pn)
					visited[pns] = struct{}{}
				}
				mn := minus(cur, i)
				mns := string(mn)
				if _, ok := visited[mns]; !ok {
					next = append(next, mn)
					visited[mns] = struct{}{}
				}
			}
		}
		queue = next
		res++
	}
	return -1
}
