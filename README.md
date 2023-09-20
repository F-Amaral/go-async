[![CI](https://github.com/f-amaral/go-async/actions/workflows/goci.yml/badge.svg)](https://github.com/f-amaral/go-async/actions/workflows/goci.yml)
[![Lint](https://github.com/f-amaral/go-async/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/f-amaral/go-async/actions/workflows/golangci-lint.yml)
[![Codecov](https://img.shields.io/codecov/c/github/f-amaral/go-async)](https://codecov.io/gh/f-amaral/go-async)


# go-async
Repository that implements goroutines abstractions, making easy to run async code in Go.

## Installation

```bash
go get github.com/f-amaral/go-async 
```

## Features
 ### Worker Pool
Creates a pool with a fixed number of workers that will execute the tasks sent to the pool.

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/f-amaral/go-async/pool"
)

func main() {
	inputSet := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	exFunc := func(i int) (int, error) {
		return i * 2, nil
	}

	p := pool.NewPool(10, exFunc)
	defer p.Close()

	result := p.Process(inputSet)
	resultBytes, _ := json.Marshal(result)
	fmt.Println(string(resultBytes))
	// Error Handling:
	fmt.Println(result.GetErrors()) 
	
}
```

