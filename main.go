package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"goreadmongo/api"
	"goreadmongo/suites"
	"net/http"
	"os"
)


func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	parser := argparse.NewParser("goreadmongo", "Reads all messages from RocketChat. For only auditing purpose ofc.")
	// Create string flag
	w := parser.Flag("w", "whole", &argparse.Options{Required: false, Help: "Get messages from all rooms"})
	r := parser.String("r", "room-id", &argparse.Options{Required: false, Help: "Room id to penetrate"})
	b := parser.String("b", "base-url", &argparse.Options{Required: true, Help: "Base URL of your RC server, example rc.iotfox.ru"})
	l := parser.Flag("l", "list-rooms", &argparse.Options{Required: false, Help: "List all rooms ids"})
	s := parser.Flag("s", "server", &argparse.Options{Required: false, Help: "Create web api server"})
	// Parse input
	err := parser.Parse(os.Args)
	suites.BASEURL = *b
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}
	// Finally print the collected string
	if *w != false && *r != ""{
		log.Debugf("W")
		bs := bson.D{{}}
		suites.FILTER = bs
		suites.GetAllMessagesByFilter(bs, *b)
	}
	if *r != "" {
		log.Debugf("R")
		bs := bson.D{{"rid", *r}}
		_ = suites.GetAllMessagesByFilter(bs, *b)
	}
	if *l != false{
		log.Debugf("L")
		_ = suites.GetAllRooms()
	}
	if *s != false{
		r := mux.NewRouter()
		sr := r.PathPrefix("/api/v1").Subrouter()
		sr.HandleFunc("", api.AllMessagesAllRooms).Methods(http.MethodGet)
		sr.HandleFunc("", api.InternalError)
		sr.HandleFunc("/{roomID}", api.AllMessagesFilterRooms).Methods(http.MethodGet)
		log.Fatal(http.ListenAndServe(":8080", r))
	}


}
