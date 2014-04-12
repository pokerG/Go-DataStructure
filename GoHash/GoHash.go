package GoHash

import (
	"math"
	"strconv"
)

var A float64 = (math.Sqrt(5) - 1) / 2

func HashDiv(k int, m int) int {
	return k % m
}

func HashMul(k int, m int) int {
	// return int(math.Floor(float64(k) * A % 1 * float64(m)))
	return int(math.Floor(math.Mod(float64(k)*A, 1)) * float64(m))
}

func HashBit(k int, m int) int {
	s := strconv.Itoa(k)
	hash := len(s)
	for i := 0; i < len(s); i++ {
		hash = (hash << 4) ^ (hash >> 28) ^ int([]byte(s)[i])
	}
	hash = int(math.Abs(float64(hash)))
	return hash % m
}
