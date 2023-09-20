package async

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

// GetError return bool when job has failed and the error
func (j *JobResult) GetError() (bool, error) {
	if j.Err != nil {
		return true, j.Err
	}
	return false, nil
}

// GetErrors return all errors from failed jobs, it uses HasError flag to minimize allocations on getting errors
func (p *ProcessResult) GetErrors() (errs []error) {
	if p.HasError {
		for _, r := range p.Results {
			if r.Err != nil {
				errs = append(errs, r.Err)
			}
		}
	}
	return errs
}
