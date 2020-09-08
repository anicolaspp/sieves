package worker

import (
	"fmt"
	"math"
)

//Data represents the Worker dataset and the communication channels
type Data struct {
	data   []int
	wid    int
	In     <-chan int
	out    chan<- int
	result chan []int
}

func NewData(wid int, in <-chan int, out chan<- int, result chan []int) *Data {
	return &Data{data: []int{}, wid: wid, In: in, out: out, result: result}
}

//Worker starts working goroutine to process a partition of the dataset
// On each iteration, each Worker filter those values in its partition that are multiple of next
// and send to the master the local next of the partition.
// If the partition has been exhausted, it sends -1 to the master
func Worker(wid int, size int, data *Data) {

	initPartition(wid, size, data)

	for {
		printMsg(wid, data.data)
		// gather next value (to be filtered) from master
		next := <-data.In
		printMsg(wid, fmt.Sprintf("Next filter: %v", next))

		// check if termination signal has been received
		if next == -1 {
			data.result <- filterInvalid(data.data)  // filter -1 to reduce communication cost
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
			data.out <- math.MaxInt32
		}
	}
}

//initPartition initializes Worker partition
func initPartition(wid int, size int, data *Data) {
	for i := wid*size + 1; i < wid*size+size+1; i++ {
		data.data = append(data.data, i)
	}
}

func printMsg(wid int, value interface{}) {
	//fmt.Printf("%v: %v\n", wid, value)
}

//GatherData sends a signal to each Worker so they send their local primes to the corresponding channel
func GatherData(chs []chan int, result chan []int) []int {
	primes := make([]int, 0)

	for i := 0; i < len(chs); i++ {
		// send signal to gather Worker data
		chs[i] <- -1

		// gather Worker primes
		primes = append(primes, <-result...)
	}

	return primes
}

//filterInvalid removes unnecessary values from the result (-1)
func filterInvalid(unfiltered []int) []int {
	primes := make([]int, 0)

	for _, v := range unfiltered {
		if v > -1 {
			primes = append(primes, v)
		}
	}
	return primes
}
