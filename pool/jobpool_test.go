package pool_test

import (
	"sort"
	"testing"

	"github.com/f-amaral/go-async/async"

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
	t.Run("Should return sorted output when created with sorting output", func(t *testing.T) {
		// arrange
		dataSet := make([]int, 100)
		for i := 0; i < 100; i++ {
			dataSet[i] = i
		}

		sut := pool.NewPool[int, int](10, func(i int) (int, error) {
			return i, assert.AnError
		}, pool.WithSortingOutput[int, int]())

		// act
		res := sut.Process(dataSet)

		// assert
		assert.Equal(t, len(dataSet), len(res.Results))
		assert.True(t, res.HasError)
		assert.True(t, sort.SliceIsSorted(res.Results, func(i, j int) bool {
			return res.Results[i].ExecIndex < res.Results[j].ExecIndex
		}))
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

// region Benchmarks
var result async.ProcessResult

func buildData() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func processWithoutErr(i int) (int, error) {
	return i * 2, nil
}

func processWithErr(i int) (int, error) {
	return i, assert.AnError
}

func BenchmarkJobPool_Process_WithoutErr(b *testing.B) {
	var processResult async.ProcessResult
	p := pool.NewPool(10, processWithoutErr)
	data := buildData()
	defer p.Close()
	for i := 0; i < b.N; i++ {
		processResult = p.Process(data)
	}
	result = processResult
}

func BenchmarkJobPool_Process_WithErr(b *testing.B) {
	var processResult async.ProcessResult
	p := pool.NewPool(10, processWithErr)
	data := buildData()
	defer p.Close()
	for i := 0; i < b.N; i++ {
		processResult = p.Process(data)
	}
	result = processResult
}

// endregion
