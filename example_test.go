package workpool

func ExampleWorkPool_struct() {
	numWorkers := 5
	outputs := make(chan int, numWorkers)

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

	pool.Run()
}

func ExampleNew() {
	numWorkers := 5
	outputs := make(chan int, numWorkers)

	worker := func(abort <-chan struct{}) bool {
		outputs <- 1
		return false
	}

	pool := New(numWorkers, worker)
	pool.Run()
	close(outputs)
}

func ExampleNewWithClose() {
	numWorkers := 5
	outputs := make(chan int, numWorkers)

	worker := func(abort <-chan struct{}) bool {
		outputs <- 1
		return false
	}
	closer := func() {
		close(outputs)
	}

	pool := NewWithClose(numWorkers, worker, closer)
	go pool.Run()
}
