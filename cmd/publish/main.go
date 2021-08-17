package main

import (
	"fmt"
	"log"
	"os"
	"play"
)

func main() {
	result, err := play.Publish(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}