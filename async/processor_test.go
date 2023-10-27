package async_test

import (
	"errors"
	"sort"
	"testing"

	"github.com/f-amaral/go-async/async"
	"github.com/stretchr/testify/assert"
)

func TestProcessResult(t *testing.T) {
	t.Run("Should get errors when jobs have errors", func(t *testing.T) {

		jobResults := buildJobResultsWithErrs()
		processResult := async.ProcessResult{
			Results:  jobResults,
			HasError: true,
		}
		errs := processResult.GetErrors()
		assert.Equal(t, len(jobResults), len(errs))
	})
	t.Run("Should get errors when jobs dont have errors", func(t *testing.T) {

		jobResults := buildJobResultsWithoutErrs()
		processResult := async.ProcessResult{
			Results:  jobResults,
			HasError: false,
		}
		errs := processResult.GetErrors()
		assert.Equal(t, 0, len(errs))
	})
	t.Run("Should return correct sorted outputs", func(t *testing.T) {
		jobResults := buildJobResultsWithoutErrs()
		processResult := async.ProcessResult{
			Results:  jobResults,
			HasError: false,
		}
		processResult.Sort()
		assert.True(t, sort.SliceIsSorted(processResult.Results, func(i, j int) bool {
			return processResult.Results[i].GetIndex() < processResult.Results[j].GetIndex()
		}))
	})
}

func TestJobResult(t *testing.T) {
	t.Run("Should return correct index when getting index", func(t *testing.T) {

		jobResult := async.JobResult{
			ExecIndex: 10,
		}
		assert.Equal(t, 10, jobResult.GetIndex())
	})
}

// region Benchmarks
var (
	errs []error
)

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

// endregion
