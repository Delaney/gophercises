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
	fmt.Println("Enter number:")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		// number := parseInt(strings.Replace(text, "\n", "", -1))
		if number, err := strconv.Atoi(text); err == nil {
			fmt.Printf("Input: %d", number)
			fmt.Println("")
			fmt.Printf("Type: %T", number)
		} else {
			fmt.Println("Enter a valid integer")
		}

	}
}
