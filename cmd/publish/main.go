package main

import (
	"fmt"
	"log"
	"os"
	"play"
)

func main() {
	p, err := play.NewPublisherFromArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.Publish())
}