package pq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/pq"
)

func testMinPQ(t *testing.T, pq pq.IPQ) {
	arr := []int{1, 2, 10, 5, 6, 8, 3}
	k := 4
	expected := 5
	for _, num := range arr {
		pq.Add(num)
		if pq.Size() > k {
			pq.DelTop()
		}
	}
	assert.Equal(t, expected, pq.Peek())
}

func testMaxPQ(t *testing.T, pq pq.IPQ) {
	arr := []int{1, 2, 10, 5, 6, 8, 3}
	k := 4
	expected := 5
	for _, num := range arr {
		pq.Add(num)
		if pq.Size() > k {
			pq.DelTop()
		}
	}
	assert.Equal(t, expected, pq.Peek())
}

func TestPQ(t *testing.T) {
	p1 := pq.NewPQ(7, pq.MIN)
	testMinPQ(t, p1)

	p2 := pq.NewPQ(7, pq.MAX)
	testMaxPQ(t, p2)
}

func TestPQ2(t *testing.T) {
	p1 := pq.NewPQ2(7, pq.MIN)
	testMinPQ(t, p1)

	p2 := pq.NewPQ2(7, pq.MAX)
	testMaxPQ(t, p2)
}
