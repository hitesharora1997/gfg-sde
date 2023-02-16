package main

import "fmt"

func main() {
	var j int
	fmt.Println("Enter the number")
	fmt.Scanf("%d", &j)

	x := 0
	for j > 0 {
		j = j / 10
		x++
	}

	fmt.Println(x)
}
