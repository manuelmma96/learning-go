/* Exercise-02.
Write a program that declares a constant called value that can be assigned to both an integer and a floating-point variable.
Assign it to an integer called i and a floating-point variable called f.
Print out i and f.
*/

package main

import "fmt"

func main() {
	const value = 30
	i := value
	var f float64 = value
	fmt.Println(i)
	fmt.Println(f)
}
