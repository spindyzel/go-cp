package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	var b []byte
	for ; n > 0; n-- {
		if fmt.Scan(&b); b[1] == '+' {
			x++
		} else {
			x--
		}
	}
	fmt.Print(x)
}
