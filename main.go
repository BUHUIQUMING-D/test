package main

import "fmt"

func main() {
	fmt.Println(Add(1, 3, 4, 5)) // output 13
}

func Add(a ...int) (result int) {
	for _, v := range a {
		result += v //   result = result + v
	}
	return result
}
