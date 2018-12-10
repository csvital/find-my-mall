package dao
 
import (
	"log"
 
	. "github.com/csvital/find_my_mall/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
 
type ShoppingMallsDAO struct {
	Server   string
	Database string
}
 
var db *mgo.Database
 
const (
	COLLECTION = "shopping_malls"
)
 
func (m *ShoppingMallsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of shopping malls
func (m *ShoppingMallsDAO) FindAll() ([]ShoppingMall, error) {
	var shopping_malls []ShoppingMall
	err := db.C(COLLECTION).Find(bson.M{}).All(&shopping_malls)
	return shopping_malls, err
}

// Find a shopping mall by its id
func (m *ShoppingMallsDAO) FindById(id string) (ShoppingMall, error) {
	var shopping_mall ShoppingMall
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&shopping_mall)
	return shopping_mall, err
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