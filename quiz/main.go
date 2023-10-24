package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	if file, err := os.Open("problems.csv"); err == nil {
		defer file.Close()

		csvReader := csv.NewReader(file)
		if problems, err := csvReader.ReadAll(); err == nil {
			var score int = 0
			for _, problem := range problems {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print(problem[0] + "= ")

				text, _ := reader.ReadString('\n')
				text = strings.Replace(text, "\n", "", -1)

				if text == problem[1] {
					score++
				}
			}

			fmt.Println("Test Complete!")
			fmt.Printf("\nScore: %d/%d", score, len(problems))
		} else {
			fmt.Println("Error reading problems")
		}

	} else {
		log.Fatal("Error while reading the file", err)
	}

}
