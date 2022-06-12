package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Footballstars struct {
	Id int
	Name string
	Position string
	Networth int
}
	var footballstars = []Footballstars{
		{1, "Messi", "Striker", 34000000000000},
		{2, "Messi", "Winger", 455500000000},
		{3, "Ronaldo", "Striker",122230033300000},
		{4, "Mbappe", "Striker", 983412210000000},
	}

	func returnAllFootballStars(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(footballstars)
		
	}

	func returnFootballStarsByName(w http.ResponseWriter, r *http.Request){
		
		vars := mux.Vars(r)
		Stars := vars["Name"]
		fstars := &[]Footballstars{}
		for _, fstar := range footballstars{
			if fstar.Name == Stars{
				*fstars = append(*fstars, fstar )
			}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(fstars)
	}

	func returnFootballStarsById(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		Starsid, err := strconv.Atoi(vars["Id"])
		if err != nil{
			fmt.Println("unable to convert to string")
		}
		for _, fottballstarsid := range footballstars{
			if fottballstarsid.Id == Starsid{
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(fottballstarsid)
			}
		}
		
	}

	func createFootballStars(w http.ResponseWriter, r *http.Request){
		var newStars Footballstars
		json.NewDecoder(r.Body).Decode(&newStars)
		footballstars = append(footballstars, newStars)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(footballstars)
		
	}

	 func removeFootballStarsbyId(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		Starsid, err := strconv.Atoi(vars["Id"])
		if err != nil{
			fmt.Println("unable to convert to string")
		}
		for k,v := range footballstars{
			if v.Id == Starsid{
				footballstars = append(footballstars[:k], footballstars[k+1:]...)
			}
		}
	 	w.WriteHeader(http.StatusOK)
		 json.NewEncoder(w).Encode(footballstars)
		
	 }

	func UpdateFootballStars(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
		Starsid, err := strconv.Atoi(vars["Id"])
		if err != nil{
			fmt.Println("unable to convert to string")
		}
		var updatedstars Footballstars
		json.NewDecoder(r.Body).Decode(&updatedstars)
		for k,v := range footballstars{
			if v.Id == Starsid{
				footballstars = append(footballstars[:k], footballstars[k+1:]...)
				footballstars = append(footballstars, updatedstars)
			}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(footballstars)
	}

func main(){
	
	router :=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/footballStars", returnAllFootballStars).Methods("GET")
	router.HandleFunc("/footballStars/stars/{Name}", returnFootballStarsByName).Methods("GET")
	router.HandleFunc("/footballStars/{Id}", returnFootballStarsById).Methods("GET")
	router.HandleFunc("/footballStars/{Id}", UpdateFootballStars).Methods("PUT")
	router.HandleFunc("/footballStars", createFootballStars).Methods("POST")
	router.HandleFunc("/footballStars/{Id}", removeFootballStarsbyId).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", router))
}