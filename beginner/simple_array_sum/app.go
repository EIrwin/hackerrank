package main

import (
	"fmt"
)

func main() {

	//grab number of arguments
	var num int
	fmt.Scanln(&num)

	//sum up input string
	sum := 0
	for i := 0; i < num; i++ {
		var val int
		fmt.Scan(&val)
		sum += val
	}

	fmt.Println(sum)
}
