package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)


func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	app := GetCli()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
