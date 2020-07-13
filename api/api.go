package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"goreadmongo/suites"
	"net/http"
)

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