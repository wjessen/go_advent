package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// read the file
	dat, err := os.ReadFile("day_one_text.txt")
	check(err)

	apples := strings.Split(string(dat), "\n")
	var total_changes = 0
	for i := 0; i < (len(apples) - 1); i++ {

		// Skip the first value
		if i == 0 {
			continue
		}

		prev_val_int, err := strconv.Atoi(apples[i-1])
		this_val_int, err := strconv.Atoi(apples[i])

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		if this_val_int > prev_val_int {
			total_changes += 1
		}

	}

	fmt.Printf("The Total Changes is: %d \n", total_changes)
}
