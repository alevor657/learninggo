package main

import "fmt"

func main() {
	sol := average([]float64{1.12341234, 2.4234, 3.24234, 4.5342534, 6.3245345})
	fmt.Println(sol)
}

func average(slice []float64) (avg float64) {
	var sum float64

	for _, val := range slice {
		sum += val
	}

	avg = sum / float64(len(slice))
	return
}

func bubbleSort(slice []int {
	for i, item := range slice {
		// TODO:
	}
}
