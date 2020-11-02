package solve0349

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// todo: helper
func arrFromJSON(input string) []int {
	var res []int
	_ = json.Unmarshal([]byte(input), &res)
	return res
}

func TestIntersection(t *testing.T) {
	assert.Equal(t, arrFromJSON("[2]"), Intersection(arrFromJSON("[1,2,2,1]"), arrFromJSON("[2,2]")))
}
