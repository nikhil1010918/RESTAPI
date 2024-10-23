package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var profiles []Profile = []Profile{}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    User   `json:"employee"`
}

func addItem(q http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	q.Header().Set("content-type", "application/json")
	profiles = append(profiles, newProfile)

	json.NewEncoder(q).Encode(profiles)

}
func getAllProfiles(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("content-Type", "application/json")
	json.NewEncoder(q).Encode(profiles)

}

func getProfile(q http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["ID"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("no profile found with specified ID"))
		return
	}
	profile := profiles[id]
	q.Header().Set("content-Type", "application/json")
	json.NewEncoder(q).Encode(profile)

}

func updateProfile(q http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["ID"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("no profile found with specified ID"))
		return
	}
	var updateProfile Profile
	json.NewDecoder(r.Body).Decode(&updateProfile)

	profiles[id] = updateProfile

	q.Header().Set("Content-type", "application/json")
	json.NewEncoder(q).Encode(updateProfile)

}

func deleteProfile(q http.ResponseWriter, r *http.Request) {

	var idParam string = mux.Vars(r)["ID"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("no profile found with specified ID"))
		return
	}

	profiles = append(profiles[:id], profiles[id+1:]...)

	q.WriteHeader(200)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/profiles", addItem).Methods("POST")

	router.HandleFunc("/profiles", getAllProfiles).Methods("GET")

	router.HandleFunc("/profiles/{ID}", getProfile).Methods("GET")

	router.HandleFunc("/profiles/{ID}", updateProfile).Methods("PUT")

	router.HandleFunc("/profiles/{ID}", deleteProfile).Methods("DELETE")

	http.ListenAndServe(":5000", router)
}
