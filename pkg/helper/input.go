package helper

import "encoding/json"

// ArrFromJSON json str -> []int
func ArrFromJSON(input string) []int {
	var res []int
	_ = json.Unmarshal([]byte(input), &res)
	return res
}

// GridFromJSON json str -> [][]int
func GridFromJSON(input string) [][]int {
	var grid [][]int
	_ = json.Unmarshal([]byte(input), &grid)
	return grid
}
