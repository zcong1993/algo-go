package solve0753

/**
@index 753
@title 打开转盘锁
@difficulty 中等
@tags breadth-first-search
@draft false
@link https://leetcode-cn.com/problems/open-the-lock/
@frontendId 752
*/
func openLock(deadends []string, target string) int {
	start := "0000"
	deadMap := make(map[string]struct{}, len(deadends))
	for _, d := range deadends {
		deadMap[d] = struct{}{}
	}
	if _, ok := deadMap[start]; ok {
		return -1
	}
	queue := make([]string, 0)
	queue = append(queue, start)
	mm := make(map[string]struct{}, 0)
	mm[start] = struct{}{}
	step := 0

	nearMap := map[byte][]byte{
		'0': {'1', '9'},
		'1': {'0', '2'},
		'2': {'1', '3'},
		'3': {'2', '4'},
		'4': {'3', '5'},
		'5': {'4', '6'},
		'6': {'5', '7'},
		'7': {'6', '8'},
		'8': {'7', '9'},
		'9': {'8', '0'},
	}

	for len(queue) > 0 {
		next := make([]string, 0)
		for _, str := range queue {
			if str == target {
				return step
			}
			bts := []byte(str)
			for i := 0; i < 4; i++ {
				for _, near := range nearMap[bts[i]] {
					origin := bts[i]
					bts[i] = near
					nextStr := string(bts)
					_, deadExists := deadMap[nextStr]
					_, mpExists := mm[nextStr]
					if !deadExists && !mpExists {
						mm[nextStr] = struct{}{}
						next = append(next, nextStr)
					}
					bts[i] = origin
				}
			}
		}
		queue = next
		step++
	}

	return -1
}

func OpenLock(deadends []string, target string) int {
	return openLock(deadends, target)
}
