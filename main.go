package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"goreadmongo/suites"
	"os"

)

func main() {
	parser := argparse.NewParser("goreadmongo", "Reads all messages from RocketChat. For only auditing purpose ofc.")
	// Create string flag
	r := parser.String("r", "room-id", &argparse.Options{Required: true, Help: "Room id to penetrate"})
	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}
	// Finally print the collected string
	suites.GetAllMessagesFromRoom(*r)
}