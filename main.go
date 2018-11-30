package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetSlaves ...
func GetSlaves(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://ori-ental.ni.corp.natinst.com/api/systems/test/slaves")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(body))
}

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/slaves", GetSlaves).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
