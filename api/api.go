package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"goreadmongo/suites"
	"net/http"
)

func StartServer(port int, allowedOringinPort int){
	r := mux.NewRouter()
	sr := r.PathPrefix("/api/v1").Subrouter()
	sr.HandleFunc("/messages", AllMessagesAllRooms).Methods(http.MethodGet)
	sr.HandleFunc("", InternalError)
	sr.HandleFunc("/ping", Ping)
	sr.HandleFunc("/rooms", Rooms)
	sr.HandleFunc("/messages/{roomID}", AllMessagesFilterRooms).Methods(http.MethodGet)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{fmt.Sprintf("http://localhost:%d",allowedOringinPort)},
		AllowCredentials: true,
	})

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",port), handler))
}
func Ping(w http.ResponseWriter, r *http.Request) {
	log.Infof("Ping")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mes,_ := json.Marshal("Pong!")
	w.Write(mes)
	return
}

func Rooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		res := suites.GetAllRooms()
		resbyte,err := json.Marshal(res)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "Internal Error"}`))
			return

		}
		w.WriteHeader(http.StatusOK)
		w.Write(resbyte)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Bad method"}`))
		return
	}
}

func AllMessagesAllRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		res := suites.GetAllMessagesByFilter(bson.D{},suites.BASEURL)
		resbyte,err := json.Marshal(res)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "Internal Error"}`))
			return

		}
		w.WriteHeader(http.StatusOK)
		w.Write(resbyte)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Bad method"}`))
		return
	}
}

func InternalError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{"message": "Internal Error"}`))
	return
}

//TODO
func AllMessagesFilterRooms(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	if val, ok := pathParams["roomID"]; ok {
		res := suites.GetAllMessagesByFilter(bson.D{{"rid",val}},suites.BASEURL)
		resbyte,err := json.Marshal(res)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "Internal Error"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resbyte)
			return

	}
}