package pool

import (
	"github.com/f-amaral/go-async/domain"
	"sync"
)

type JobFn[I any, O any] func(I) (O, error)

type jobPool[I any, O any] struct {
	workers int
	jobFn   JobFn[I, O]
}

func NewPool[I any, O any](workers int, job JobFn[I, O]) domain.Processor[I, O] {
	return &jobPool[I, O]{
		workers: workers,
		jobFn:   job,
	}
}

func (s *jobPool[I, O]) Process(inputs []I) domain.ProcessResult {
	inputSize := len(inputs)

	var wg sync.WaitGroup
	jobs := make(chan I, inputSize)
	results := make(chan domain.JobResult, inputSize)

	for w := 0; w < min(inputSize, s.workers); w++ {
		wg.Add(1)
		go func(jobs <-chan I, results chan<- domain.JobResult) {
			defer wg.Done()
			for _, input := range inputs {
				output, err := s.jobFn(input)
				results <- domain.JobResult{
					Input:  input,
					Output: output,
					Err:    err,
				}
			}
		}(jobs, results)
	}

	for _, input := range inputs {
		jobs <- input
	}
	close(jobs)

	defer func() {
		wg.Wait()
		close(results)
	}()

	output := domain.ProcessResult{
		Results:  make([]domain.JobResult, 0),
		HasError: false,
	}

	for range inputs {
		result := <-results
		if result.Err != nil {
			output.HasError = true
		}
		output.Results = append(output.Results, result)
	}

	return output
}
