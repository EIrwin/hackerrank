package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
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
	scanner.Split(bufio.ScanWords)

	var sum uint64
	for scanner.Scan() {
		valStr := scanner.Text()
		val,err := strconv.ParseUint(valStr,10,64)
		if err != nil {
			fmt.Println("error")
			log.Fatalln(err)
		}
		fmt.Println(val)
		sum += val
	}
	fmt.Println(sum)

}
