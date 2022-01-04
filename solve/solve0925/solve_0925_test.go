package solve0925_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0925"
)

func TestConstructFromPrePost(t *testing.T) {
	pre := helper.ArrFromJSON(`[4,2,1,3,7,6,9]`)
	pos := helper.ArrFromJSON(`[1,3,2,6,9,7,4]`)
	expected := "4,2,1,#,#,3,#,#,7,6,#,#,9,#,#"

	te := solve0925.ConstructFromPrePost(pre, pos)
	assert.Equal(t, expected, tree.Serialize(te))
}
