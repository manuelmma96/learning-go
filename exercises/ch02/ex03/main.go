/* Exercise 03
Write a program with three variables, one named b of type byte, one named smallI of type int32, and one named bigI of type uint64.
Assign each variable to the maximum legal value for its type; then add 1 to each variable. Print out their values
*/
package main
import "fmt"

func main() {

	/*Solution #1 */
	// With constant expressions
    // var b byte =  255 + 1 // Compile time overflow
    // var smallI int32 = 2147483647 + 1 // Compile time overflow
    // var bigI uint64 = 18446744073709551615 + 1 // Compile time overflow
    // fmt.Println(b, smallI, bigI)


	/*Solution #2 */
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	// With variables
	b = b + 1 //Run-time wrap around (Integer Overflow)
	smallI = smallI + 1 //Run-time wrap around (Integer Overflow)
	bigI = bigI + 1 //Run-time wrap around (Integer Overflow)

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
