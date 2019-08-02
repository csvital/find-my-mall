package dao

import (
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
	// COLLECTION is the name of the collection in MongoDB
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
	var shoppingMalls []ShoppingMall
	err := db.C(COLLECTION).Find(bson.M{}).All(&shoppingMalls)
	return shoppingMalls, err
}

// FindByID Find a shopping mall by its id
func (m *ShoppingMallsDAO) FindByID(id string) (ShoppingMall, error) {
	var shoppingMalls ShoppingMall
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&shoppingMalls)
	return shoppingMalls, err
}

// FindByQuery Find a shopping mall by query
func (m *ShoppingMallsDAO) FindByQuery(city, score, sortField string, shopsToFind []string) ([]ShoppingMall, error) {
	var shoppingMalls []ShoppingMall
	var internal []interface{}
	if len(shopsToFind) > 0 {
		for _, element := range shopsToFind {
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
	err := db.C(COLLECTION).Find(queryToFind).Sort(sortField).All(&shoppingMalls)
	return shoppingMalls, err
}

// Insert a shopping mall into database
func (m *ShoppingMallsDAO) Insert(shoppingMalls ShoppingMall) error {
	err := db.C(COLLECTION).Insert(&shoppingMalls)
	return err
}

// Delete an existing shopping mall
func (m *ShoppingMallsDAO) Delete(shoppingMalls ShoppingMall) error {
	err := db.C(COLLECTION).Remove(&shoppingMalls)
	return err
}

// Update an existing shopping mall
func (m *ShoppingMallsDAO) Update(shoppingMalls ShoppingMall) error {
	err := db.C(COLLECTION).UpdateId(shoppingMalls.ID, &shoppingMalls)
	return err
}
