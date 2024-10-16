package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y)
	z = x - (x * y / 100)
	fmt.Print(z)
}
