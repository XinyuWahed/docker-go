package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("hello world")
	if err := health(); err != nil {
		log.Fatal("failed to start health handler", err)
	}
}


