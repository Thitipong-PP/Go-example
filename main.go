// package main

// import (
// 	"fmt"

// 	testpackage "github.com/Thitipong-PP/Go-example/test_package"

// 	"github.com/google/uuid"
// )

// func main () {
// 	const fullName string = "Thitipong"
// 	var age int = 1000

// 	var lastName = "Phuangphet"
// 	var io = 18

// 	x := "is God"
// 	t := 9

// 	println((fullName + " " + lastName + " " + x))
// 	println((age + io + t))
// 	// testpackage.Say("Hello Guy, It me Thitipong Phuangphet")
// 	testpackage.Say(uuid.NewString())

// 	// For loop
// 	// for i:=0; i<10; i++ {
// 	// 	println(i)
// 	// }

// 	// While loop
// 	// i := 0
// 	// for i<10 {
// 	// 	println(i)
// 	// 	i++;
// 	// }

// 	// Do-While loop
// 	// i := 0
// 	// for {
// 	// 	println(i)
// 	// 	i++;
// 	// 	if i >= 10 {
// 	// 		break
// 	// 	}
// 	// }

// 	var arr [3]int
// 	arr[0] = 1
// 	arr[1] = 9
// 	arr[2] = 7

// 	var sld []int

// 	sld = arr[:]
// 	sld = append(sld, 8)
// 	sld = append(sld, 91821, 8123, 12312312, 12838641)

// 	fmt.Println(arr)
// 	fmt.Println(sld)
// }

package main

import (
	"fmt"
)

func main () {
	sayHello("Thitipong", add(1, 2))
	sayHello("Phuangphet", 4)
	sayHello("Is god", 1)
}

func sayHello (say string, round int) {
	for round > 0 {
		fmt.Println(say)
		round--
	}
}

func add (a int, b int) int {
	return a + b
}