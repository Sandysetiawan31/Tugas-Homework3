package changoroutine

type adder interface {
	AddCounter(int64)
}

func TryChanGoroutine(intChan chan<- int, add adder) {
	// Add the logic to complete the tests in here

	for i := 1; i <= 4; i++ {
		intChan <- 11
	}
	intChan <- 3

	for j := 1; j <= 50; j++ {
		go add.AddCounter(1)
	}

}
