package sort_test

import (
	"github.com/f-amaral/go-async/sort"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

type sortable struct {
	Data  string
	Index int
}

func (s *sortable) GetIndex() int {
	return s.Index
}

func NewUnsortedArray(size int) []*sortable {
	var arr []*sortable
	for i := 0; i < size; i++ {
		arr = append(arr, &sortable{
			Data:  "AnyData",
			Index: i,
		})
	}

	for i := 0; i < size; i++ {
		shuffleIndex := rand.Intn(size)
		arr[i], arr[shuffleIndex] = arr[shuffleIndex], arr[i]
	}

	return arr
}

func TestMerge(t *testing.T) {
	// Arrage
	unsortedArr := NewUnsortedArray(10)

	// Act
	sortedArr := sort.Merge(unsortedArr)

	for i, element := range sortedArr {
		// Assert
		assert.Equal(t, i, element.Index)
	}
}
