package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	NewScanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Width of the rhombus? ")
	NewScanner.Scan()
	_MaxLength, _ := strconv.ParseInt(NewScanner.Text(), 10, 64)
	MaxLength := int(_MaxLength)

	if MaxLength % 2 == 0 {
		MaxLength += 1
	}

	for i := 1; i <= MaxLength; i++ {
		OddNumber := 2*i - 1
		// 2 * 1 - 1 = 1

		put_space := PutSpace(MaxLength, OddNumber)
		// 5 - 1 / 2 = 4
		// 6 - 1 / 2 = 2

		if OddNumber > 0 && OddNumber <= MaxLength {
			for i := 1; i <= put_space; i++ {
				fmt.Print(" ")
			}
			for i := 1; i <= OddNumber; i++ {
				fmt.Print(".")
			}
			fmt.Println()
		}
	}
	for i := MaxLength; i >= 0; i-- {
		OddNumber := i - 2
		put_space := PutSpace(MaxLength , OddNumber)

		if 0 < OddNumber && OddNumber % 2 != 0 {
			for a := 0; a < put_space; a++ {
				fmt.Print(" ")
			}
			for d := 0; d < OddNumber; d++ {
				fmt.Print(".")
			}
			fmt.Println()
		}
	}
}

func PutSpace(MaxLength, OddNumber int) int {
	return (MaxLength - OddNumber) / 2
}
