package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/flaviogf/go-doubler"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("at least 1 argument is required")
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatalln("the argument must be an integer")
	}

	fmt.Printf("%d\n", doubler.Double(int32(n)))
}
