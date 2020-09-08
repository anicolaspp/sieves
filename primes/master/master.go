package master

import "math"

//Master starts Master process
func Master(workers int, chs []chan int, In chan int) {
	next := 2

	for {
		broadcastMin(chs, next)

		min, hasMin := agreeOnNextMin(workers, In)

		if !hasMin {
			break
		}

		next = min
	}
}

//broadcastMin broadcast the newest selected, global min
func broadcastMin(chs []chan int, next int) {
	for i := 0; i < len(chs); i++ {
		chs[i] <- next
	}
}

//agreeOnNextMin collects the local min of each partition/worker and select to global min
func agreeOnNextMin(workers int, In chan int) (int, bool) {
	min := math.MaxInt32

	for i := 0; i < workers; i++ {
		localMin := <-In

		min = _min(min, localMin)
	}

	if min == math.MaxInt32 {
		return -1, false
	}

	return min, true
}



func _min(a, b int) int  {
	if a <= b {
		return a
	}

	return b
}
