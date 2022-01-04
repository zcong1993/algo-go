package solve1050_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve1050"
)

func TestBstFromPreorder(t *testing.T) {
	input := helper.ArrFromJSON(`[8,5,1,7,10,12]`)
	expected := "8,5,1,#,#,7,#,#,10,#,12,#,#"

	assert.Equal(t, expected, tree.Serialize(solve1050.BstFromPreorder(input)))
}
