package testpackage

import "fmt"

// This function like private in OOP cannot use in other package
func pro (x string) {
	fmt.Println(x)
}

// This function like public in OOP
func Say (tell string) {
	// fmt.Println(tell)
	pro(tell) // Can use in package
}