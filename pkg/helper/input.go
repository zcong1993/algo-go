package helper

import "encoding/json"

// ArrFromJSON json str -> []int
func ArrFromJSON(input string) []int {
	var res []int
	_ = json.Unmarshal([]byte(input), &res)
	return res
}

// ArrFromJSONString json str -> []string
func ArrFromJSONString(input string) []string {
	var res []string
	_ = json.Unmarshal([]byte(input), &res)
	return res
}

// GridFromJSON json str -> [][]int
func GridFromJSON(input string) [][]int {
	var grid [][]int
	_ = json.Unmarshal([]byte(input), &grid)
	return grid
}

// GridFromJSONString json str -> [][]string
func GridFromJSONString(input string) [][]string {
	var grid [][]string
	_ = json.Unmarshal([]byte(input), &grid)
	return grid
}

// StringBoardFromJSON json str -> [][]int
func StringBoardFromJSON(input string) [][]string {
	var board [][]string
	_ = json.Unmarshal([]byte(input), &board)
	return board
}

// CharBoardFromJSON json str -> [][]byte
func CharBoardFromJSON(input string) [][]byte {
	board := StringBoardFromJSON(input)
	res := make([][]byte, len(board))
	for j, line := range board {
		l := make([]byte, len(line))
		for i, str := range line {
			l[i] = str[0]
		}
		res[j] = l
	}
	return res
}
