package handlers

import (
	"HelpNow/models"
	"HelpNow/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//SaveMeetings : ""
func SaveMeetings(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	meetings := new(models.Meetings)

	err := json.NewDecoder(r.Body).Decode(&meetings)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = services.SaveMeetings(meetings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	m := make(map[string]interface{})
	m["data"] = meetings
	ResponseV1(w, "Successfully registered", m)
}

//GetMeeting : ""
func GetMeeting(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	uniqueID := vars["uniqueID"]

	data, err := services.GetMeeting(uniqueID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	m := make(map[string]interface{})
	m["meeting details"] = data
	ResponseV1(w, "Participant Meeting Details", m)
}

//GetAllMeetings : ""
func GetAllMeetings(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	start := r.URL.Query().Get("start")

	end := r.URL.Query().Get("end")

	data, err := services.GetAllMeetings(start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	m := make(map[string]interface{})
	m["Meeting Information"] = data
	ResponseV1(w, "List of All Meeting", m)
}

//GetAllUserMeetingsByEmail : ""
func GetAllUserMeetingsByEmail(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	participant := r.URL.Query().Get("participant")

	data, err := services.GetAllUserMeetingsByEmail(participant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	m := make(map[string]interface{})
	m["Meeting Information"] = data
	ResponseV1(w, "List of Participant Meeting", m)
}

//AllowCors :
func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}

//ResponseV1 : ""
func ResponseV1(w http.ResponseWriter, msg string, data map[string]interface{}) {
	response := new(models.Response)
	response.StatusCode = 200
	response.Message = msg
	response.Data = data
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataB)
}
