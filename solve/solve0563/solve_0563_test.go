package solve0563

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestFindTilt(t *testing.T) {
	assert.Equal(t, 1, FindTilt(tree.Deserialize("1,2,#,#,3,#,#")))
	assert.Equal(t, 15, FindTilt(tree.Deserialize("4,2,3,#,#,5,#,#,9,#,7,#,#")))
}
