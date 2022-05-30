package main

import "fmt"

func main() {
	i, j := 42, 2701

	// print out normal value
	fmt.Println(i)
	fmt.Println(j)

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)

}
