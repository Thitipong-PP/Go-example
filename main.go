package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main () {
	print("Hello\n")
	fmt.Println("Thitipong Phuangphet")
	fmt.Println(uuid.New().String())
}