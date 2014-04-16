package GoSort

import (
	"math/rand"
	"time"
)

const smallDataSize = 512

const mediumDataSize = 16384

const largeDataSize = 65536

func testDataSorted() []int {
	input := make([]int, largeDataSize)
	for i := 0; i < largeDataSize; i++ {
		input[i] = i
	}
	return input
}

func testDataReversed() []int {
	input := make([]int, largeDataSize)
	for i := 0; i < largeDataSize; i++ {
		input[i] = largeDataSize - i - 1
	}
	return input
}

func testDataRandom() []int {
	rand.Seed(time.Now().Unix())
	input := make([]int, largeDataSize)
	for i := 0; i < largeDataSize; i++ {
		input[i] = rand.Int()
	}
	return input
}
