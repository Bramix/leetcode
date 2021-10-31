package main

import (
	"fmt"
	"math"
)

func main() {

	var bits []int = countBits(5)

	for i := range bits {
		fmt.Println(bits[i])
	}
}

func countBits(n int) []int {
	var array []int
	for i := 0; i < n + 1; i++ {
		if i == 0 {
			array = append(array, 0)
			continue
		}
		var logOfNumber = math.Log2(float64(i))
		var logOfNumberInt64 = int64(logOfNumber)
		if float64(logOfNumberInt64) == logOfNumber {
			array = append(array, 1)
			continue
		}
		var remnant = i - int(math.Pow(float64(2), float64(logOfNumberInt64)))
		array = append(array, array[remnant] + 1)
	}
	return array
}
