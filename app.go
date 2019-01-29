package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	. "github.com/csvital/find_my_mall/config"
	. "github.com/csvital/find_my_mall/dao"
	. "github.com/csvital/find_my_mall/models"
	"github.com/gorilla/mux"
)

var config = Config{}
var dao = ShoppingMallsDAO{}

// AllShoppingMalls GET list of shopping_malls
func AllShoppingMalls(w http.ResponseWriter, r *http.Request) {
	shoppingMall, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, shoppingMall)
}

// FindShoppingMall GET a shopping manll by its ID
func FindShoppingMall(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shoppingMall, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Shopping Mall ID")
		return
	}
	respondWithJSON(w, http.StatusOK, shoppingMall)
}

// CreateShoppingMall POST a new shopping mall
func CreateShoppingMall(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var shoppingMall ShoppingMall
	if err := json.NewDecoder(r.Body).Decode(&shoppingMall); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	shoppingMall.ID = bson.NewObjectId()
	if err := dao.Insert(shoppingMall); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, shoppingMall)
}

// UpdateShoppingMall PUT update an existing shopping mall
func UpdateShoppingMall(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var shoppingMall ShoppingMall
	if err := json.NewDecoder(r.Body).Decode(&shoppingMall); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(shoppingMall); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteShoppingMall DELETE an existing shopping mall
func DeleteShoppingMall(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	shoppingMall, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := dao.Delete(shoppingMall); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/shoppingMalls", AllShoppingMalls).Methods("GET")
	r.HandleFunc("/shoppingMalls", CreateShoppingMall).Methods("POST")
	r.HandleFunc("/shoppingMalls", UpdateShoppingMall).Methods("PUT")
	r.HandleFunc("/shoppingMalls/{id}", DeleteShoppingMall).Methods("DELETE")
	r.HandleFunc("/shoppingMalls/{id}", FindShoppingMall).Methods("GET")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully connected")
	}
}
