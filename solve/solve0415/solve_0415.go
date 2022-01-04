package solve0415

import (
	"strconv"
)

/**
 * @index 415
 * @title 字符串相加
 * @difficulty 简单
 * @tags math,string,simulation
 * @draft false
 * @link https://leetcode-cn.com/problems/add-strings/
 * @frontendId 415
 */

func AddStrings(num1 string, num2 string) string {
	add := 0
	i, j := len(num1)-1, len(num2)-1
	res := ""
	for i >= 0 || j >= 0 || add > 0 {
		n1, n2 := 0, 0
		r := add
		if i >= 0 {
			n1 = int(num1[i] - '0')
			r += n1
			i--
		}

		if j >= 0 {
			n2 = int(num2[j] - '0')
			r += n2
			j--
		}
		add = r / 10
		res = strconv.Itoa(r%10) + res
	}

	return res
}
