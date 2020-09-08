package primes

import "fmt"

//workerData represents the worker dataset and the communication channels
type workerData struct {
	data   []int
	wid    int
	In     <-chan int
	out    chan<- int
	result chan []int
}

//worker starts working goroutine to process a partition of the dataset
// On each iteration, each worker filter those values in its partition that are multiple of next
// and send to the master the local next of the partition.
// If the partition has been exhausted, it sends -1 to the master
func worker(wid int, size int, data *workerData) {

	initPartition(wid, size, data)

	for {
		printMsg(wid, data.data)
		// gather next value to be filtered from master
		next := <-data.In
		printMsg(wid, fmt.Sprintf("Next filter: %v", next))

		// check if termination signal has been received
		if next == -1 {
			data.result <- data.data
			return
		}

		// filter partition using next value
		for i, v := range data.data {
			if v != next && v%next == 0 {
				data.data[i] = -1
			}
		}

		// select the next local min
		hastNext := false
		for _, v := range data.data {
			if v > next {
				data.out <- v
				hastNext = true
				break
			}
		}

		// if there is not next local min, send signal of partition completed
		if !hastNext {
			data.out <- -1
		}
	}

}

//initPartition initializes worker partition
func initPartition(wid int, size int, data *workerData) {
	for i := wid*size + 1; i < wid*size+size+1; i++ {
		data.data = append(data.data, i)
	}
}

func printMsg(wid int, value interface{}) {
	//fmt.Printf("%v: %v\n", wid, value)
}
