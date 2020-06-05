package main

import (
	"fmt"
	"go/types"
)

var a int

func main() {
	v := []int{1, 2, 3}

	types.NewSlice()
	for i, v := range v {
		j := i
		x := v
		fmt.Println(&j, &x, j, x)
	}

	a = 10
	fmt.Printf("%p %d\n", &a, a)
	a = 20
	fmt.Printf("%p %d\n", &a, a)
}
