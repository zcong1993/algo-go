package solve1013_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/solve/solve1013"
)

func testFib(t *testing.T, fib func(n int) int) {
	seq := []int{0, 1, 1, 2, 3, 5, 8}
	for i, s := range seq {
		assert.Equal(t, s, fib(i))
	}
}

func TestFib(t *testing.T) {
	testFib(t, solve1013.Fib)
}

func TestFib2(t *testing.T) {
	testFib(t, solve1013.Fib2)
}

func TestFib3(t *testing.T) {
	testFib(t, solve1013.Fib3)
}

func TestFib4(t *testing.T) {
	testFib(t, solve1013.Fib4)
}
