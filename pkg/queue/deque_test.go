package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/algo-go/pkg/queue"
)

func TestDeQueue(t *testing.T) {
	q := queue.NewDeQueue()
	_, err := q.PeekLast()
	assert.Equal(t, queue.ErrNotExists, err)
	_, err = q.PollLast()
	assert.Equal(t, queue.ErrNotExists, err)
	_, err = q.PeekFirst()
	assert.Equal(t, queue.ErrNotExists, err)
	_, err = q.PollFirst()
	assert.Equal(t, queue.ErrNotExists, err)
	assert.Equal(t, 0, q.Size())
	q.AddFirst(1)
	q.AddFirst(0)
	q.AddFirst(2)
	assert.Equal(t, 3, q.Size())
	num, err := q.PeekFirst()
	assert.Nil(t, err)
	assert.Equal(t, num, 2)
	num, err = q.PeekLast()
	assert.Nil(t, err)
	assert.Equal(t, num, 1)
	num, err = q.PollFirst()
	assert.Nil(t, err)
	assert.Equal(t, num, 2)
	assert.Equal(t, 2, q.Size())
	num, err = q.PollLast()
	assert.Nil(t, err)
	assert.Equal(t, num, 1)
	assert.Equal(t, 1, q.Size())
	num, err = q.PeekFirst()
	assert.Nil(t, err)
	assert.Equal(t, num, 0)
	num, err = q.PeekLast()
	assert.Nil(t, err)
	assert.Equal(t, num, 0)
	num, err = q.PollFirst()
	assert.Nil(t, err)
	assert.Equal(t, num, 0)
	assert.Equal(t, 0, q.Size())
}
