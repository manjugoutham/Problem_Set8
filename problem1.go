/* Read LearnerData.txt file. Print the complete file read. Enhance the program to read one line at a time and print out the same.
Also find out how many valid email-ids are present in this file. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("LearnerData.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	filedata, err := os.ReadFile("LearnerData.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Complete file contents:")
	fmt.Println(string(filedata))

	fmt.Println("\n Contents line by line:")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	countValidEmails := 0

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		if emailRegex.MatchString(scanner.Text()) {
			countValidEmails++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nValid email IDs: %d\n", countValidEmails)
}
