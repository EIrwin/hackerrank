package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {

	//number of rocks
	var num int
	fmt.Scanln(&num)

	reader := bufio.NewReader(os.Stdin)

	//global count
	counts := make(map[string]int)

	//read in rocks
	for i := 1; i <= num ;i ++ {
		bytes,_,_ := reader.ReadLine()

		//read composition
		rock := string(bytes)

		//local count
		m := make(map[string]bool,len(rock))
		for _,composition := range rock {
			c := string(composition)
			m[c] = true
		}

		for key,_ := range m {
			counts[key]++
		}
	}

	gems := 0
	for _,val := range counts {
		if(val == num){
			gems++
		}
	}
	fmt.Println(gems)
}