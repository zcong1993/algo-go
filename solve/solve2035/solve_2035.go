package solve2035

/**
 * @index 2035
 * @title 统计子岛屿
 * @difficulty 中等
 * @tags depth-first-search,breadth-first-search,union-find,array,matrix
 * @draft false
 * @link https://leetcode-cn.com/problems/count-sub-islands/
 * @frontendId 1905
 */

func CountSubIslands(grid1 [][]int, grid2 [][]int) int {
	res := 0
	m, n := len(grid1), len(grid1[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid2[i][j] == 0 {
			return
		}
		grid2[i][j] = 0
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	// grid1 为水, grid2 为陆地, 不可能出现子岛屿
	// 排除掉
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}
