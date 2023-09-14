package domain

// Processor is a generic interface to process asynchronously a list of inputs.
type Processor[I any, O any] interface {
	Process([]I) ProcessResult
	Close()
}

// JobResult is the result of a single job.
type JobResult struct {
	Input  any
	Output any
	Err    error
}

// ProcessResult is the result of a list of jobs, the hasError flag indicates if any job has failed.
type ProcessResult struct {
	Results  []JobResult
	HasError bool
}
