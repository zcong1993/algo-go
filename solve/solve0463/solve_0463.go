package solve0463

/**
 * @index 463
 * @title 岛屿的周长
 * @difficulty 简单
 * @tags depth-first-search,breadth-first-search,array,matrix
 * @draft false
 * @link https://leetcode-cn.com/problems/island-perimeter/
 * @frontendId 463
 */

func IslandPerimeter(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	// 上下左右遍历偏移量矩阵
	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, 1, -1}
	ans := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// index 越界或者不是陆地, return
		if i < 0 || i > m-1 || j < 0 || j > n-1 || grid[i][j] != 1 {
			return
		}
		// 标记防止重复遍历
		grid[i][j] = 2
		// 遍历上下左右相邻节点
		for k := 0; k < 4; k++ {
			nextI, nextJ := i+dx[k], j+dy[k]
			// 相邻节点为, 越界或者海洋时, 周长+1
			if nextI < 0 || nextI > m-1 || nextJ < 0 || nextJ > n-1 || grid[nextI][nextJ] == 0 {
				ans++
				continue
			}
			dfs(nextI, nextJ)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				// 只有一个岛屿, 直接 return
				return ans
			}
		}
	}
	// 万一没有岛屿
	return ans
}
