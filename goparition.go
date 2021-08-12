package main

import (
	"fmt"
)

func main() {
	var bigList = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	const partitionSize = 3
	for idxRange := range Partition(len(bigList), partitionSize) {
		bulkOperation(bigList[idxRange.Low:idxRange.High])
	}
}

type IdxRange struct {
	Low, High int
}

func Partition(collectionLen, partitionSize int) chan IdxRange {
	c := make(chan IdxRange)
	if partitionSize <= 0 {
		close(c)
		return c
	}

	go func() {
		numFullPartitions := collectionLen / partitionSize
		var i int
		for ; i < numFullPartitions; i++ {
			c <- IdxRange{Low: i * partitionSize, High: (i + 1) * partitionSize}
		}

		if collectionLen%partitionSize != 0 { // left over
			c <- IdxRange{Low: i * partitionSize, High: collectionLen}
		}

		close(c)
	}()
	return c
}

func bulkOperation(x []int) {
	fmt.Println(x)
}
