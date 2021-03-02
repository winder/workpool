package workpool

import "fmt"

func ExampleWorkPool_struct() {
	numWorkers := 2
	outputs := make(chan int)

	worker := func(abort <-chan struct{}) bool {
		outputs <- 1
		return false
	}

	closer := func() {
		close(outputs)
	}

	pool := &WorkPool{
		Handler: worker,
		Workers: numWorkers,
		Close:   closer,
	}

	go pool.Run()
	for out := range outputs {
		fmt.Println(out)
	}
	// Output: 1
	// 1
}

func ExampleNew() {
	numWorkers := 3
	outputs := make(chan int)

	worker := func(abort <-chan struct{}) bool {
		outputs <- 1
		return false
	}

	pool := New(numWorkers, worker)
	go func() {
		pool.Run()
		close(outputs)
	}()

	for out := range outputs {
		fmt.Println(out)
	}
	// Output: 1
	// 1
	// 1
}

func ExampleNewWithClose() {
	numWorkers := 5
	outputs := make(chan int)

	worker := func(abort <-chan struct{}) bool {
		outputs <- 1
		return false
	}

	closer := func() {
		close(outputs)
	}

	pool := NewWithClose(numWorkers, worker, closer)
	go pool.Run()

	for out := range outputs {
		fmt.Println(out)
	}
	// Output: 1
	// 1
	// 1
	// 1
	// 1
}
