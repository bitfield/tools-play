package main

import (
	"fmt"
	"log"
	"os"
	"play"
)

func main() {
	result, err := play.Publish(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}