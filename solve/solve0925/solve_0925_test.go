package solve0925

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestConstructFromPrePost(t *testing.T) {
	root := ConstructFromPrePost(helper.ArrFromJSON("[1,2,4,5,3,6,7]"), helper.ArrFromJSON("[4,5,2,6,7,3,1]"))
	assert.Equal(t, "1,2,4,#,#,5,#,#,3,6,#,#,7,#,#", tree.Serialize(root))
}
