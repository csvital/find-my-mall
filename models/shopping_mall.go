package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type ShoppingMall struct {
	ID			bson.ObjectId `bson:"_id" json:"id"`
	Name		string `bson:"name" json:"name"`
	District	string `bson:"district" json:"district"`
	City		string `bson:"city" json:"city"`
}