package routes

import (
	"HelpNow/handlers"

	"github.com/gorilla/mux"
)

//MeetingRoutes :"user routes"
func MeetingRoutes(r *mux.Router) {
	r.HandleFunc("/meetings", handlers.SaveMeetings).Methods("POST")
	r.HandleFunc("/meetings/{uniqueID}", handlers.GetMeeting).Methods("GET")
	r.HandleFunc("/meetings", handlers.GetAllMeetings).Methods("GET")
	r.HandleFunc("/meeting", handlers.GetAllUserMeetingsByEmail).Methods("GET")
}

//ThreadSafe: ""
func ThreadSafeRoutes(r *mux.Router) {
	go func() {
		MeetingRoutes(r)
	}()

	go func() {
		MeetingRoutes(r)
	}()
}
