package solve0054

/**
 * @index 54
 * @title 螺旋矩阵
 * @difficulty 中等
 * @tags array,matrix,simulation
 * @draft false
 * @link https://leetcode-cn.com/problems/spiral-matrix/
 * @frontendId 54
 */

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	num := len(matrix) * len(matrix[0])
	res := make([]int, 0, num)

	t, b := 0, len(matrix)-1
	l, r := 0, len(matrix[0])-1
	// 收缩边界
	for len(res) < num {
		for i := l; i <= r && len(res) < num; i++ {
			res = append(res, matrix[t][i])
		}
		t++
		for i := t; i <= b && len(res) < num; i++ {
			res = append(res, matrix[i][r])
		}
		r--
		for i := r; i >= l && len(res) < num; i-- {
			res = append(res, matrix[b][i])
		}
		b--
		for i := b; i >= t && len(res) < num; i-- {
			res = append(res, matrix[i][l])
		}
		l++
	}
	return res
}
