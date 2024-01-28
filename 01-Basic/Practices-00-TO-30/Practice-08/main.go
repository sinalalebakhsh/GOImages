package main

import (
	"fmt"
)

func main() {
	maxLength := 21
	for i := 0; i < maxLength; i++ {
		OddNumber := 2 * i - 1
		PutSpace := (maxLength - OddNumber) / 2 
		/* 
			11 - 1  = 10 / 2  = 5
			11 - 3  = 8  / 2  = 4
			11 - 5  = 6  / 2  = 3
			11 - 7  = 4  / 2  = 2
			11 - 9  = 2  / 2  = 1
			11 - 11 = 0  / 2  = 0
		*/
		
		if 0 < OddNumber && OddNumber <= maxLength {
			for a := 0; a < PutSpace; a++ {
				fmt.Print(" ")
			}
			for d := 0; d < OddNumber; d++ {
				fmt.Print("*")
			}
			fmt.Println()
		} 
	}
	
	for i := maxLength; i >= 0; i-- {
		OddNumber := i - 2 
		PutSpace := (maxLength - OddNumber) / 2 

		if 0 < OddNumber && OddNumber % 2 != 0 {
			for a := 0; a < PutSpace; a++ {
				fmt.Print(" ")
			}
			for d := 0; d < OddNumber; d++ {
				fmt.Print("*")
			}
			fmt.Println()	
		}
	}
}
