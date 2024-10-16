package main

import "fmt"

func main() {
	var x, y, z float64
	fmt.Print("masukan berat badan (kg): ")
	fmt.Scan(&x)
	fmt.Print("masukan tinggi badan (m): ")
	fmt.Scan(&y)
	z = x / (y * y)
	fmt.Printf("BMI anda: %.2f", z)
}
