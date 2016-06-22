package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

const (
	ALPHABET_COUNT = 26
)

func main()  {
	m := make(map[string]bool,ALPHABET_COUNT)

	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')

	count := 0
	for _,char := range input {

		c := strings.ToLower(strings.TrimSpace(string(char)))

		if !m[c] && c != " " && c != ""{
			m[c] = true
			count++
		}
	}

	if count == 26 {
		fmt.Println("pangram")
	}else {
		fmt.Println("not pangram")
	}

}
