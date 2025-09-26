/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-09-26
 * @fileoverview volume of sphere
 */

package main

import "fmt"

func main() {

	// state question
	fmt.Println("What is the volume of a sphere with a radius of 4 cm?")

	// state formula
	fmt.Println("V = 4 / 3 * Pi * r * r * r")

	// put in known information
	fmt.Println("V = 4 / 3 * 3.14 * 4 * 4 * 4")

	// solve the problem
	fmt.Println("V = " + fmt.Sprint(4.0/3.0) + " * 3.14 * 4 * 4 * 4")

	fmt.Println("V = " + fmt.Sprint(1.33*3.14) + " * 4 * 4 * 4")

	fmt.Println("V = " + fmt.Sprint(4.18*4) + " * 4 * 4")
	
	fmt.Println("V = " + fmt.Sprint(16.72*4) + " * 4")

	fmt.Println("V = " + fmt.Sprint(66.88*4))

	// final statement
	fmt.Println("The volume of a sphere with a radius of 4 cm is ~267.52 cm.")

	fmt.Println("\nDone.")

}
