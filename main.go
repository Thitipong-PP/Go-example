package main

import (
	testpackage "github.com/Thitipong-PP/Go-example/test_package" // Package name testpackage

	"github.com/google/uuid"
)

func main () {
	testpackage.Say("Hello Guy, It me Thitipong Phuangphet")
	testpackage.Say(uuid.NewString())
}