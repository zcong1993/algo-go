package solve1050

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/pkg/tree"
)

func TestBstFromPreorder(t *testing.T) {
	assert.Equal(t, "8,5,1,#,#,7,#,#,10,#,12,#,#", tree.Serialize(BstFromPreorder(helper.ArrFromJSON("[8,5,1,7,10,12]"))))
}

func TestBstFromPreorder2(t *testing.T) {
	assert.Equal(t, "8,5,1,#,#,7,#,#,10,#,12,#,#", tree.Serialize(BstFromPreorder2(helper.ArrFromJSON("[8,5,1,7,10,12]"))))
}
