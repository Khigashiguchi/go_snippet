package main

import (
	"./rand"
	"fmt"
)

func main() {
	var length int
	fmt.Printf("Please input the length which you want to generate.\n")
	fmt.Scan(&length)
	fmt.Println("Output:", rand.String(length))
}
