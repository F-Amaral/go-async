package pool_test

import (
	"testing"

	"github.com/f-amaral/go-async/pool"
	"github.com/stretchr/testify/assert"
)

func TestJobPool_Process(t *testing.T) {
	t.Run("Should process all inputs without error", func(t *testing.T) {
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
	t.Run("should process all inputs and return error", func(t *testing.T) {
		// arrange
		dataSet := make([]int, 100)
		for i := 0; i < 100; i++ {
			dataSet[i] = i
		}

		sut := pool.NewPool[int, int](10, func(i int) (int, error) {
			return i, assert.AnError
		})

		// act
		res := sut.Process(dataSet)

		// assert
		assert.Equal(t, len(dataSet), len(res.Results))
		assert.True(t, res.HasError)
	})
}

func TestJobPool_Close(t *testing.T) {
	t.Run("Should close the job channel", func(t *testing.T) {
		// arrange
		sut := pool.NewPool[int, int](10, func(i int) (int, error) {
			return i, nil
		})

		// act
		sut.Close()

		// assert
		assert.Panics(t, func() {
			sut.Process([]int{1, 2, 3})
		})
	})
}
