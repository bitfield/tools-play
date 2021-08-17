package main

import (
	"fmt"
	"os"
	"play"
)

func main() {
	result := play.Publish(os.Args[1:])
	fmt.Println(result)
}