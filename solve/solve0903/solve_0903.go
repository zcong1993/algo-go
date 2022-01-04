package solve0903

import "math/rand"

/**
 * @index 903
 * @title 用 Rand7() 实现 Rand10()
 * @difficulty 中等
 * @tags math,rejection-sampling,probability-and-statistics,randomized
 * @draft false
 * @link https://leetcode-cn.com/problems/implement-rand10-using-rand7/
 * @frontendId 470
 */

func rand7() int {
	return rand.Intn(7) + 1
}

// https://github.com/zcong1993/kanban/issues/80
// https://leetcode-cn.com/problems/implement-rand10-using-rand7/solution/mo-neng-gou-zao-fa-du-li-sui-ji-shi-jian-9xpz
// https://leetcode-cn.com/problems/implement-rand10-using-rand7/solution/xiang-xi-si-lu-ji-you-hua-si-lu-fen-xi-zhu-xing-ji

// 大到小直接舍弃部分就行

// (randX() - 1)*Y + randY() -> [1, X * Y]
// (rand7() - 1) * 7 + rand7() -> [1, 49] -> [1, 40] -> [1, 10]

// rand3 -> rand5
// (rand3() - 1) * 3 + rand3() -> [1, 9] -> [1, 5]

// rand7 -> rand100
// rand3 = rand7 舍弃 7 -> rand6() % 3 + 1
// rand49 = (rand7() - 1) * 7 + rand7()
// rand147 = (rand49() - 1) * 3 + rand3() -> 舍弃 > 100 -> rand100

func Rand10() int {
	res := (rand7()-1)*7 + rand7()
	if res > 40 {
		return Rand10()
	}
	return 1 + res%10
}

// rand7 -> rand100

func rand3() int {
	res := rand7()

	if res > 6 {
		return rand3()
	}

	return res%3 + 1
}

func rand49() int {
	return (rand7()-1)*7 + rand7()
}

func Rand100() int {
	res := (rand49()-1)*3 + rand3()
	if res > 100 {
		return Rand100()
	}
	return res
}
