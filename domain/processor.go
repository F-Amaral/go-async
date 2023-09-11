package domain

type JobResult struct {
	Input  any
	Output any
	Err    error
}

type ProcessResult struct {
	Results  []JobResult
	HasError bool
}

type Processor[I any, O any] interface {
	Process([]I) ProcessResult
}
