package models

import "gopkg.in/mgo.v2/bson"

type McdonaldsItem struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Title       string 	  `bson:"title" json:"title"`
	Image       string 	  `bson:"image" json:"image"`
	Price       int    	  `bson:"price" json:"price"`
	Description string 	  `bson:"description" json:"description"`
}
