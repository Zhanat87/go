package apis

import (
	"github.com/go-ozzo/ozzo-routing"
	"os"
	"gopkg.in/mgo.v2"
	"github.com/Zhanat87/go/models"
)

func getMcdonaldsItems() []models.McdonaldsItem {
	session, err := mgo.Dial(os.Getenv("MONGODB_DSN"))
	if err != nil {
		panic("ERROR db connect: " + err.Error())
	}
	itemsCollection := session.DB("mcdonalds").C("items")
	var items []models.McdonaldsItem
	err = itemsCollection.Find(nil).All(&items)
	if err != nil {
		panic("ERROR find items: " + err.Error())
	}
	return items
}

func Mcdonalds() routing.Handler {
	return func(c *routing.Context) error {
		return c.Write(getMcdonaldsItems())
	}
}
