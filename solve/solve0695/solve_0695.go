package solve0695

/**
 * @index 695
 * @title 岛屿的最大面积
 * @difficulty 中等
 * @tags depth-first-search,breadth-first-search,union-find,array,matrix
 * @draft false
 * @link https://leetcode-cn.com/problems/max-area-of-island/
 * @frontendId 695
 */

func MaxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		if grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				r := dfs(i, j)
				if r > res {
					res = r
				}
			}
		}
	}
	return res
}
