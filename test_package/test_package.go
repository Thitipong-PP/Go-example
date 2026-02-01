package testpackage

import (
	"fmt"

	"github.com/Thitipong-PP/Go-example/test_package/internal/idk"
)

// This function like private in OOP cannot use in other package
func pro (x string) {
	fmt.Println(x)
}

// This function like public in OOP
func Say (tell string) {
	// fmt.Println(tell)
	idk.Test()
	pro(tell) // Can use in package
}