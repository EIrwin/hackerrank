package main

import (
	"fmt"
)

func main()  {
	var input string
	fmt.Scanln(&input)

	if(checkPalindrome(input)){
		fmt.Println("YES")
	}else {
		fmt.Println("NO")
	}
}

func checkPalindrome(s string) (bool) {

	counts := make(map[string]int, 26)

	for _, ch := range s {
		c := string(ch)
		counts[c]++
	}

	oddChars := getOddCharacterCount(counts)

	return oddChars <= 1
}

func getOddCharacterCount(c map[string]int) (int)  {

	count := 0
	for _,val := range c {
		if(!isEven(val)){
			count++
		}
	}

	return  count
}

func isEven(num int) (bool)  {
	return num % 2 == 0
}