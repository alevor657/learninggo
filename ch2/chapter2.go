package main

import "fmt"

func main() {
	// exercise1()
	// exercise2()
	// exercise3()
	// exercise4([]float64{2.1234123, 1.12341234, 2.12341234, 1.12341234, 3.0, 1000.0})
	fizzBuzz()
}

func exercise1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func exercise2() {
	func() {
		i := 0
	l1:
		if i >= 10 {
			return
		}
		fmt.Println(i)
		i++
		goto l1
	}()
}

func exercise3() {
	var arr [10]int

	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	fmt.Println(arr)
}

func exercise4(slice []float64) {
	var sum float64

	for _, nr := range slice {
		sum += nr
	}

	avg := sum / float64(len(slice))
	fmt.Println(avg)
}

func fizzBuzz() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
