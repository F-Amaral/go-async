package pool

import (
	"github.com/f-amaral/go-async/domain"
)

type JobFn[I any, O any] func(I) (O, error)

type jobParams[I any] struct {
	input     I
	resultsCh chan domain.JobResult
}

type jobPool[I any, O any] struct {
	jobCh chan<- jobParams[I]
}

func NewPool[I any, O any](workers int, jobFunc JobFn[I, O]) domain.Processor[I, O] {
	jobs := make(chan jobParams[I], workers)

	for w := 0; w < workers; w++ {
		go func() {
			for params := range jobs {
				output, err := jobFunc(params.input)
				params.resultsCh <- domain.JobResult{
					Input:  params.input,
					Output: output,
					Err:    err,
				}
			}
		}()
	}

	return &jobPool[I, O]{
		jobCh: jobs,
	}
}

func (s *jobPool[I, O]) Process(inputs []I) domain.ProcessResult {
	inputSize := len(inputs)

	results := make(chan domain.JobResult, inputSize)

	for _, input := range inputs {
		s.jobCh <- jobParams[I]{
			input:     input,
			resultsCh: results,
		}
	}

	defer func() {
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

func (s *jobPool[I, O]) Close() {
	close(s.jobCh)
}
