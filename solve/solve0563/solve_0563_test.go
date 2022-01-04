package solve0563_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0563"
)

func TestFindTilt(t *testing.T) {
	input := tree.Deserialize("21,7,1,3,#,#,3,#,#,1,#,#,14,2,#,#,2,#,#")
	expected := 9

	assert.Equal(t, expected, solve0563.FindTilt(input))
}
