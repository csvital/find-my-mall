package models

import "gopkg.in/mgo.v2/bson"

// ShoppingMall represents the structure of the shopping mall, we uss bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type ShoppingMall struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	District   string        `bson:"district" json:"district"`
	City       string        `bson:"city" json:"city"`
	ImageLink  string        `bson:"imageLink" json:"imageLink"`
	DetailPage string        `bson:"detailPage" json:"detailPage"`
	Score      string        `bson:"score" json:"score"`
	ShopList   []Shop        `bson:"shops" json:"shops"`
	CafesList  []Cafe        `bson:"cafes" json:"cafes"`
}

// Shop is the structure of the shops inside a shopping mall
type Shop struct {
	Logo    string `bson:"logo" json:"logo"`
	Magaza  string `bson:"magaza" json:"magaza"`
	Kat     string `bson:"kat" json:"kat"`
	Telefon string `bson:"telefon" json:"telefon"`
}

// Cafe is the structure of the cafe inside a shopping mall
type Cafe struct {
	Logo    string `bson:"logo" json:"logo"`
	Magaza  string `bson:"magaza" json:"magaza"`
	Kat     string `bson:"kat" json:"kat"`
	Telefon string `bson:"telefon" json:"telefon"`
}
