package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scanf("%d %v\n", &a, &b)

	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
