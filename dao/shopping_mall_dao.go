package dao

import (
	"fmt"
	"log"

	. "github.com/works-forces/find-my-mall/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ShoppingMallsDAO is the DAO structure
type ShoppingMallsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "shopping_malls"
)

// Connect connects to the MongoDB
func (m *ShoppingMallsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// FindAll Find list of shopping malls
func (m *ShoppingMallsDAO) FindAll() ([]ShoppingMall, error) {
	var shopping_malls []ShoppingMall
	err := db.C(COLLECTION).Find(bson.M{}).All(&shopping_malls)
	return shopping_malls, err
}

// FindById Find a shopping mall by its id
func (m *ShoppingMallsDAO) FindById(id string) (ShoppingMall, error) {
	var shopping_mall ShoppingMall
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&shopping_mall)
	return shopping_mall, err
}

// FindByQuery Find a shopping mall by query
func (m *ShoppingMallsDAO) FindByQuery(city, score string, shopsToFind []string) ([]ShoppingMall, error) {
	var shopping_malls []ShoppingMall
	var internal []interface{}
	if len(shopsToFind) > 0 {
		for _, element := range shopsToFind {
			println(element)
			internal = append(internal, bson.M{"shops": bson.M{"$elemMatch": bson.M{"magaza": element}}})
		}
	} else {
		internal = append(internal, bson.M{})
	}
	if city != "" {
		internal = append(internal, bson.M{"city": city})
	}
	if score == "" {
		score = "0,0"
	}
	queryToFind := bson.M{"score": bson.M{"$gte": score}, "$and": internal}
	fmt.Println(queryToFind)
	err := db.C(COLLECTION).Find(queryToFind).All(&shopping_malls)
	return shopping_malls, err
}

// Insert a shopping mall into database
func (m *ShoppingMallsDAO) Insert(shopping_mall ShoppingMall) error {
	err := db.C(COLLECTION).Insert(&shopping_mall)
	return err
}

// Delete an existing shopping mall
func (m *ShoppingMallsDAO) Delete(shopping_mall ShoppingMall) error {
	err := db.C(COLLECTION).Remove(&shopping_mall)
	return err
}

// Update an existing shopping mall
func (m *ShoppingMallsDAO) Update(shopping_mall ShoppingMall) error {
	err := db.C(COLLECTION).UpdateId(shopping_mall.ID, &shopping_mall)
	return err
}
