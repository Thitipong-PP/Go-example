package main

import (
	testpackage "github.com/Thitipong-PP/Go-example/test_package"

	"github.com/google/uuid"
)

func main () {
	var fullName string = "Thitipong"
	var age int = 1000

	var lastName = "Phuangphet"
	var io = 18


	x := "is God"
	t := 9

	println((fullName + " " + lastName + " " + x))
	println((age + io + t))
	// testpackage.Say("Hello Guy, It me Thitipong Phuangphet")
	testpackage.Say(uuid.NewString())
}