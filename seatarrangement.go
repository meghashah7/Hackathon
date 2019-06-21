/*
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here
package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func main() {
	
	var A, B, k int
	fmt.Scanf("%d", &A)
	for k = 1; k <= A; k++ {
		fmt.Scanf("%d", &B)
		i, j := solution(B)
		fmt.Println(i, j)
	}

}
func solution(A int) (int, string) {
	// write your code in Go 1.4
	if A < 109 && A > 0 {
		switch A % 12 {
		case 0:
			return A - 11, "WS"
		case 11:
			return A - 9, "MS"
		case 10:
			return A - 7, "AS"
		case 9:
			return A - 5, "AS"
		case 8:
			return A - 3, "MS"
		case 7:
			return A - 1, "WS"
		case 6:
			return A + 1, "WS"
		case 5:
			return A + 3, "MS"
		case 4:
			return A + 5, "AS"
		case 3:
			return A + 7, "AS"
		case 2:
			return A + 9, "MS"
		case 1:
			return A + 11, "WS"
		default:
			return 0, "ERROR"

		}
	} else {
		return 0, "ERROR"
	}
}

