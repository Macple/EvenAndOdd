package main

import "fmt"

func main() {
	iSlice := make([]int, 10)

	for i := range iSlice {
		isEven(i)
	}
}

func isEven(i int) {

	if i%2 == 0 {
		fmt.Printf("%v is even\n", i)
	} else {
		fmt.Printf("%v is odd\n", i)

	}

}
