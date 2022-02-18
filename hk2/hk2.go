package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		i uint64 = 4
		d        = 4.0
		s        = "HackerRank "
	)

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	// Read and save an integer, doubl2
	//e, and String to your variables.
	// var ints string
	scanner.Scan()
	i2, _ := strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	d2, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	s2 := scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + i2)

	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f", d+d2)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Printf("\n%s%s", s, s2)

}
