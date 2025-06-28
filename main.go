// main.go
package main

import (
	"fmt"
	"learn-circleci/internal/mathutil"
)

func main() {
	fmt.Println("Hello, Modular Go!")

	sum := mathutil.Add(3, 5)
	sum = mathutil.Sub(sum, 2)
	fmt.Printf("3 + 5 = %d\n", sum)
}
