package solve0105

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
)

// todo: helper
func arrFromJSON(input string) []int {
	var res []int
	_ = json.Unmarshal([]byte(input), &res)
	return res
}

func TestBuildTree(t *testing.T) {
	expected := "1,2,4,#,#,5,#,#,3,6,#,#,7,#,#"
	root := BuildTree(arrFromJSON("[1,2,4,5,3,6,7]"), arrFromJSON("[4,2,5,1,6,3,7]"))
	assert.Equal(t, expected, tree.Serialize(root))
}
