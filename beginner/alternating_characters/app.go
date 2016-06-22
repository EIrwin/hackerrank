package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	//number of lines
	var num int
	fmt.Scanln(&num)

	//holds repeat counts
	deletions := make([]int,num)

	//iterate line-by-line
	count := 0
	for i := 0; i < num; i ++ {

		//read in line
		reader := bufio.NewReader(os.Stdin)
		bytes,_,_ := reader.ReadLine()
		line := string(bytes)

		//count repeats
		var prev string
		for _,val := range line {
			c := string(val)
			if prev == c {
				count++
			}
			prev = c
		}

		deletions[i] = count

		//reset count
		count = 0
	}

	print(deletions)
}

func print(d []int)  {
	for _,val := range d {
		fmt.Println(val)
	}
}
