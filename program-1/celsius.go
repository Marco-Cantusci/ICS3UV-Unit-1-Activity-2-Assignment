/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-09-26
 * @fileoverview Celsius to Fahrenheit
 */

package main

import "fmt"

func main() {

	// state question
	fmt.Println("What is 34 Celsius in Fahrenheit?")

	// formula
	fmt.Println("F = 9 / 5 * C + 32")

	// put information into the variables
	fmt.Println("F = 9 / 5 * 34 + 32")

	// solve
	fmt.Println("F = " + fmt.Sprint(9.0/5.0) + " * 34 + 32")

	fmt.Println("F = " + fmt.Sprint(1.8*34) + " + 32")

	fmt.Println("F = " + fmt.Sprint(61.2+32))

	// final statement
	fmt.Println("34 Celsius is 93.2 Fahrenheit.")

	fmt.Println("\nDone.")

}
