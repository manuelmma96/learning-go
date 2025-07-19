/* Chapter 3 Exercise 1: Subslices */

/* Write a program that defines a variable named greetings of type slice of strings with the following values:
 "Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт". 
 Create a sub-slice containing the first two values.
 Create a second subslice with the second, third, and fourth values
 Create a third subslice with the fourth and fifth values.
Print out all four slices. */

package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	x := greetings[:2]
	y := greetings[1:4]
	z := greetings[3:]

	fmt.Println("Original slice:", greetings)
	fmt.Println("Slice 1", x)
	fmt.Println("Slice 2", y)
	fmt.Println("Slice 3", z)
}