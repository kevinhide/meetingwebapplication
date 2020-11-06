package main

import (
	"HelpNow/daos"
	"HelpNow/handlers"
	"HelpNow/routes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	db := daos.GetDB()
	db.Session.Close()
	fmt.Println("listenandservice 8084")
	fmt.Println("started", time.Now())
	r := mux.NewRouter()
	r.Use(handlers.AllowCors)
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})

	routes.ThreadSafeRoutes(r)

	log.Fatal(http.ListenAndServe(":8084", r))
}
