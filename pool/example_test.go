package pool_test

import (
	"fmt"
	"sync"

	"github.com/f-amaral/go-async/pool"
)

// ExampleJobPool_ProcessOneDataSet demonstrates how to use the JobPool to process one dataset.
func ExampleJobPool_Process_OneDataSet() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create a new worker pool, with 5 workers, that will process the input dataset.
	p := pool.NewPool[int, int](5, func(i int) (int, error) {
		multiplied := i * 2
		fmt.Println("multiplied", i, "by 2 and got", multiplied)
		return i * 2, nil
	})

	// It is important to close the pool when you are done using it, otherwise you can have a leaking channel.
	defer p.Close()

	p.Process(numbers)
}

// ExampleJobPool_Process demonstrates how to use the JobPool to more than one dataset with the same shared workers.
func ExampleJobPool_Process_MultipleInputsInSameWorkerPool() {
	odds := []int{1, 3, 5, 7, 9}
	evens := []int{2, 4, 6, 8, 10}

	p := pool.NewPool[int, int](5, func(i int) (int, error) {
		multiplied := i * 2
		fmt.Println("multiplied", i, "by 2 and got", multiplied)
		return i * 2, nil
	})

	// It is important to close the pool when you are done using it, otherwise you can have a leaking channel.
	defer p.Close()

	// Waitgroup is used to process multiple datasets in parallel, but is not required when using the JobPool.
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		p.Process(odds)
	}()
	go func() {
		defer wg.Done()
		p.Process(evens)
	}()
	wg.Wait()
}
