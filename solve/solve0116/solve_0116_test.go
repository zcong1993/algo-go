package solve0116_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/tree"
	"github.com/zcong1993/algo-go/solve/solve0116"
)

func convertTree(root *tree.TreeNode) *solve0116.Node {
	if root == nil {
		return nil
	}
	rc := &solve0116.Node{Val: root.Val}
	rc.Left = convertTree(root.Left)
	rc.Right = convertTree(root.Right)
	return rc
}

func testConnect(t *testing.T, connect func(root *solve0116.Node) *solve0116.Node) {
	input := tree.Deserialize("1,2,4,#,#,5,#,#,3,6,#,#,7,#,#")
	root := convertTree(input)

	nextMap := map[int]int{
		2: 3,
		4: 5,
		5: 6,
		6: 7,
	}

	connect(root)

	var inorder func(node *solve0116.Node)
	inorder = func(node *solve0116.Node) {
		if node == nil {
			return
		}
		if next, ok := nextMap[node.Val]; ok {
			assert.Equal(t, next, node.Next.Val)
		}
	}
	inorder(root)
}

func TestConnect(t *testing.T) {
	testConnect(t, solve0116.Connect)
}

func TestConnect2(t *testing.T) {
	testConnect(t, solve0116.Connect2)
}
