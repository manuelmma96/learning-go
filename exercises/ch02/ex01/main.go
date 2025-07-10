/* Exercise-01. 
Write a program that declares an integer variable called i with the value 20.
Assign 20 i to a floating-point variable named f.
Print out i and f 
*/

package main

import "fmt"

func main() {
	var i byte = 20
	var f float64 = float64(i)
	fmt.Println(i, f)
}
