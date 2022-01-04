package solve0652_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/helper"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0652"
)

func toStrings(ts []*tree.TreeNode) []string {
	res := make([]string, len(ts))
	for i, t := range ts {
		res[i] = tree.Serialize(t)
	}

	return res
}

func TestFindDuplicateSubtrees(t *testing.T) {
	testCases := []struct {
		input    *tree.TreeNode
		expected []string
	}{
		{
			input:    tree.Deserialize("0,0,0,#,#,#,0,#,0,#,0,#,#"),
			expected: helper.ArrFromJSONString(`["0,#,#"]`),
		},
		{
			input:    tree.Deserialize("1,2,4,#,#,#,3,2,4,#,#,#,4,#,#"),
			expected: helper.ArrFromJSONString(`["2,4,#,#,#", "4,#,#"]`),
		},
	}

	for _, c := range testCases {
		res := solve0652.FindDuplicateSubtrees(c.input)
		assert.ElementsMatch(t, c.expected, toStrings(res))
	}
}
