package main

import "fmt"

func main() {
	var x, y, z float64
	fmt.Scan(&x, &y)
	z = x * y * y
	fmt.Printf("%.0f", z)
}
