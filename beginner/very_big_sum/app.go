package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main()  {

	//grab number of inputs
	var num int
	fmt.Scanln(&num)

	//error check input range
	if num < 0 || num > 10 {
		log.Fatalln("Invalid range")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputStr := scanner.Text()
		input := strings.Split(inputStr," ")
		sum,_ := sum(input)
		fmt.Println(sum)
	}
}

func sum(input []string) (uint64,error)  {
	var sum uint64
	for _,value  := range input {
		val,err := strconv.ParseUint(value,10,64)
		sum += val
		if err != nil {
			log.Println(err)
		}
	}
	return sum,nil
}
