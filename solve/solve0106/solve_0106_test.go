package solve0106_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0106"
)

func TestBuildTree(t *testing.T) {
	ino := helper.ArrFromJSON(`[1,2,3,4,6,7,9]`)
	pos := helper.ArrFromJSON(`[1,3,2,6,9,7,4]`)
	expected := "4,2,1,#,#,3,#,#,7,6,#,#,9,#,#"

	assert.Equal(t, expected, tree.Serialize(solve0106.BuildTree(ino, pos)))
}
