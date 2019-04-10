package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	. "github.com/works-forces/find-my-mall/models"
	"gopkg.in/mgo.v2/bson"
)

var dummyID string

func CreateDummyShoppingMall() {
	dummyMall := &ShoppingMall{ID: bson.NewObjectId(), Name: "dummyMall", District: "dummyDistrict", City: "dummyCity"}
	jsonStr, err := json.Marshal(dummyMall)
	if err != nil {
		log.Fatal(err)
	}
	url := "http://localhost:3000/shoppingMalls"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 201 {
		log.Fatal("Cannot create dummy connection. Status code: ", res.Status)
	}
}

func GetAllDummyShoppingMalls() {
	url := "http://localhost:3000/shoppingMalls"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		log.Fatal("Cannot delete dummy connection. Status codesss: ", res.Status)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var data []map[string]interface{}
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	dummyID = data[0]["id"].(string)
}

func DeleteDummyShoppingMall() {
	url := "http://localhost:3000/shoppingMalls/" + dummyID
	req, _ := http.NewRequest("DELETE", url, nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		log.Fatal("Cannot delete dummy connection. Status code: ", res.Status)
	}
}

func PutDummyShoppingMall() {
	dummyMall := &ShoppingMall{ID: bson.ObjectIdHex(dummyID), Name: "dummyMallChanged", District: "dummyDistrictChanged", City: "dummyCityChanged"}
	jsonStr, err := json.Marshal(dummyMall)
	if err != nil {
		log.Fatal(err)
	}
	url := "http://localhost:3000/shoppingMalls"
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		log.Fatal("Cannot create dummy connection. Status code: ", res.Status)
	}
}

func FindDummyShoppingMall() {
	url := "http://localhost:3000/shoppingMalls/" + dummyID
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != 200 {
		log.Fatal("Cannot delete dummy connection. Status code: ", res.Status)
	}
}
func TestAppRoutes(t *testing.T) {
	CreateDummyShoppingMall()
	GetAllDummyShoppingMalls()
	PutDummyShoppingMall()
	FindDummyShoppingMall()
	DeleteDummyShoppingMall()
}
