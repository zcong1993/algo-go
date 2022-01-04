package solve1380

/**
 * @index 1380
 * @title 统计封闭岛屿的数目
 * @difficulty 中等
 * @tags depth-first-search,breadth-first-search,union-find,array,matrix
 * @draft false
 * @link https://leetcode-cn.com/problems/number-of-closed-islands/
 * @frontendId 1254
 */

func ClosedIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 1 {
			return
		}
		grid[i][j] = 1
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	// 把边缘陆地全部淹没
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for i := 0; i < n; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}
