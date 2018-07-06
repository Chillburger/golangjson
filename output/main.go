package main

import (
	"flag"
	"fmt"

	"github.com/dahilb/golangjson/jsonaggre"
)

func main() {
	// process flags
	// --data /path/to/json/data
	dataPtr := flag.String("data", "", "path to json data")
	flag.Parse()
	// print datapath for validation
	fmt.Println("data path: " + *dataPtr)

	//Process json input
	var Output = jsonaggre.ProcessJson(*dataPtr)
	fmt.Println("test: " + Output.Name)
}
