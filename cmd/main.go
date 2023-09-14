package main

import (
	"encoding/json"
	"fmt"

	"github.com/f-amaral/go-async/pool"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create a new worker pool, with 5 workers, that will process the input dataset.
	p := pool.NewPool[int, int](5, func(i int) (int, error) {
		return i * 2, nil
	})

	// It is important to close the pool when you are done using it, otherwise you can have a leaking channel.
	defer p.Close()

	res := p.Process(numbers)
	b, _ := json.Marshal(res)
	fmt.Println(string(b))
}
