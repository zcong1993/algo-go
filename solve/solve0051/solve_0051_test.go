package solve0051_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/solve/solve0051"
)

func TestSolveNQueens(t *testing.T) {
	assert.Equal(t, helper.GridFromJSONString(`[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]`), solve0051.SolveNQueens(4))
	assert.Equal(t, helper.GridFromJSONString(`[[".Q....","...Q..",".....Q","Q.....","..Q...","....Q."],["..Q...",".....Q",".Q....","....Q.","Q.....","...Q.."],["...Q..","Q.....","....Q.",".Q....",".....Q","..Q..."],["....Q.","..Q...","Q.....",".....Q","...Q..",".Q...."]]`), solve0051.SolveNQueens(6))
}
