package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	if file, err := os.Open("problems.csv"); err == nil {
		defer file.Close()

		csvReader := csv.NewReader(file)
		if problems, err := csvReader.ReadAll(); err == nil {
			for _, problem := range problems {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print(problem[0] + "= ")

				text, _ := reader.ReadString('\n')
				text = strings.Replace(text, "\n", "", -1)

				if number, err := strconv.Atoi(text); err == nil {
					fmt.Printf("Input: %d", number)
					fmt.Println("")
					fmt.Printf("Answer: %s", problem[1])
				} else {
					fmt.Println("Enter a valid integer")
				}
			}
		} else {
			fmt.Println("Error reading problems")
		}

	} else {
		log.Fatal("Error while reading the file", err)
	}

}
