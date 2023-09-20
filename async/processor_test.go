package async_test

import (
	"errors"
	"testing"

	"github.com/f-amaral/go-async/async"
	"github.com/stretchr/testify/assert"
)

func TestJobResult_GetError_WhenError(t *testing.T) {
	jobResult := async.JobResult{
		Input:  1,
		Output: 2,
		Err:    assert.AnError,
	}
	ok, err := jobResult.GetError()
	assert.True(t, ok)
	assert.Equal(t, assert.AnError, err)
}

func TestJobResult_GetError_WhenNoError(t *testing.T) {
	jobResult := async.JobResult{
		Input:  1,
		Output: 2,
		Err:    nil,
	}
	ok, err := jobResult.GetError()
	assert.False(t, ok)
	assert.Nil(t, err)
}

func TestProcessResult_GetErrors_WhenJobsHaveErrors(t *testing.T) {
	jobResults := buildJobResultsWithErrs()
	processResult := async.ProcessResult{
		Results:  jobResults,
		HasError: true,
	}
	errs := processResult.GetErrors()
	assert.Equal(t, len(jobResults), len(errs))
}

func TestProcessResult_GetErrors_WhenJobsDontHaveErrors(t *testing.T) {
	jobResults := buildJobResultsWithoutErrs()
	processResult := async.ProcessResult{
		Results:  jobResults,
		HasError: false,
	}
	errs := processResult.GetErrors()
	assert.Equal(t, 0, len(errs))
}

// region Benchmarks
var (
	errs []error
	err  error
)

func BenchmarkJobResult_GetError_WhenError(b *testing.B) {
	jobResult := buildJobResultWithErr()
	var result error
	for i := 0; i < b.N; i++ {
		_, result = jobResult.GetError()
	}
	err = result
}

func BenchmarkJobResult_GetError_WhenNoError(b *testing.B) {
	jobResult := buildJobResultWithoutErr()
	var result error
	for i := 0; i < b.N; i++ {
		_, result = jobResult.GetError()
	}
	err = result
}

func BenchmarkProcessResult_GetErrors(b *testing.B) {
	var processResult async.ProcessResult
	var results []error
	processResult.Results = buildJobResultsWithErrs()
	for i := 0; i < b.N; i++ {
		results = processResult.GetErrors()
	}
	errs = results
}

func BenchmarkProcessResult_GetErrors_WhenNoErrors(b *testing.B) {
	var processResult async.ProcessResult
	var results []error
	processResult.Results = buildJobResultsWithoutErrs()
	for i := 0; i < b.N; i++ {
		results = processResult.GetErrors()
	}
	errs = results
}

// endregion

// region Test Support

func buildJobResultsWithErrs() []async.JobResult {
	var jobResults []async.JobResult
	for i := 0; i < 100; i++ {
		jobResults = append(jobResults, async.JobResult{
			Input:  i,
			Output: i * 2,
			Err:    errors.New("error"),
		})
	}
	return jobResults
}

func buildJobResultsWithoutErrs() []async.JobResult {
	var jobResults []async.JobResult
	for i := 0; i < 100; i++ {
		jobResults = append(jobResults, async.JobResult{
			Input:  i,
			Output: i * 2,
			Err:    nil,
		})
	}
	return jobResults
}

func buildJobResultWithoutErr() async.JobResult {
	return async.JobResult{
		Input:  1,
		Output: 2,
		Err:    nil,
	}
}

func buildJobResultWithErr() async.JobResult {
	return async.JobResult{
		Input:  1,
		Output: 2,
		Err:    assert.AnError,
	}
}

// endregion
