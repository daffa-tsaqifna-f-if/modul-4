package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

func main() {
	var ax, ay, bx, by, cx, cy float64
	fmt.Print("masukan koordinat titik A: ")
	fmt.Scan(&ax, &ay)
	fmt.Print("masukan koordinat titik B: ")
	fmt.Scan(&bx, &by)
	fmt.Print("masukan koordinat titik C: ")
	fmt.Scan(&cx, &cy)
	AB := distance(ax, ay, bx, by)
	BC := distance(bx, by, cx, cy)
	CA := distance(cx, cy, ax, ay)
	maxSide := math.Max(AB, math.Max(BC, CA))
	fmt.Printf("%.2f\n", maxSide)
}
