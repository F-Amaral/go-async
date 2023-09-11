package pool_test

import (
	"github.com/f-amaral/go-async/pool"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJobPool_Process(t *testing.T) {
	t.Run("should process all inputs", func(t *testing.T) {
		// arrange
		dataSet := make([]int, 100)
		for i := 0; i < 100; i++ {
			dataSet[i] = i
		}

		sut := pool.NewPool[int, int](10, func(i int) (int, error) {
			return i, nil
		})

		// act
		res := sut.Process(dataSet)

		// assert
		assert.Equal(t, len(dataSet), len(res.Results))
		assert.False(t, res.HasError)
	})
}
