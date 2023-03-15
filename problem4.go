/*
Read the Cricket_Players_Stats.csv file.
a. Print all the rows
b. Print all the rows except the header row
*/
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("Cricket_Players_Stats.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Print the all rows :")
	for _, row := range rows {
		fmt.Println(row)
	}

	fmt.Println("Print the all rows except the header row : ")
	for i, row := range rows {
		if i == 0 {
			continue
		}
		fmt.Println(row)
	}
}
