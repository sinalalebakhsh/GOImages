package main

import (
	"fmt"
	"sort"
)

func main() {

	var NewSlice []string = []string{"Sina", "Mina", "Aryamehr", "Arya"}

	var NewPoint = &NewSlice[1]

	sort.Strings(NewSlice)

	fmt.Println(*NewPoint)
}
