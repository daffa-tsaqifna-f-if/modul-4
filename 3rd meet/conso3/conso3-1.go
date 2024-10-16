package main

import "fmt"

func main() {
	var s, h, m int
	fmt.Scan(&s)
	h = s / 3600
	m = (s % 3600) / 60
	s = s % 60
	fmt.Print(h, " jam ", m, " menit ", s, " dekit")
}
