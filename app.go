package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("-----------------------------------")
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1) // replace the newline character so we can cast 15 to integer

	factor, _ := strconv.ParseInt(text, 0, 0)
	fmt.Println(factorial(uint(factor)))
}

func factorial(x uint) (fac uint) {
	if x == 0 {
		return 1
	}
	fac = x * factorial(x-1)
	return
	
}
