package solve0105_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0105"
)

func TestBuildTree(t *testing.T) {
	pre := helper.ArrFromJSON(`[4,2,1,3,7,6,9]`)
	ino := helper.ArrFromJSON(`[1,2,3,4,6,7,9]`)
	expected := "4,2,1,#,#,3,#,#,7,6,#,#,9,#,#"

	assert.Equal(t, expected, tree.Serialize(solve0105.BuildTree(pre, ino)))
}
