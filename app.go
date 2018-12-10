package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/csvital/find_my_mall/config"
	. "github.com/csvital/find_my_mall/dao"
	. "github.com/csvital/find_my_mall/models"
)

var config = Config{}
var dao = ShoppingMallsDAO{}

// GET list of shopping_malls
func AllShoppingMalls(w http.ResponseWriter, r *http.Request) {
	shopping_malls, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, shopping_malls)
}

// GET a shopping manll by its ID
func FindShoppingMall(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shopping_mall, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Shopping Mall ID")
		return
	}
	respondWithJson(w, http.StatusOK, shopping_mall)
}

// POST a new shopping mall
func CreateShoppingMall(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var shopping_mall ShoppingMall
	if err := json.NewDecoder(r.Body).Decode(&shopping_mall); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	shopping_mall.ID = bson.NewObjectId()
	if err := dao.Insert(shopping_mall); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, shopping_mall)
}

// PUT update an existing shopping mall
func UpdateShoppingMall(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var shopping_mall ShoppingMall
	if err := json.NewDecoder(r.Body).Decode(&shopping_mall); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(shopping_mall); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing shopping mall
func DeleteShoppingMall(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var shopping_mall ShoppingMall
	if err := json.NewDecoder(r.Body).Decode(&shopping_mall); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(shopping_mall); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
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
	r.HandleFunc("/shoppingMalls", DeleteShoppingMall).Methods("DELETE")
	r.HandleFunc("/shoppingMalls/{id}", FindShoppingMall).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}