package main

import (
	config "Package/Generics/pkg/config"
	query "Package/Generics/pkg/query"
	//model "Package/Generics/pkg/model"
	"fmt"
)

func main() {
	db,_ := config.MongoConnection("Books") // Create instance of Database
	//collection := query.NewMongoCollection(db,"book") // Create instance of Collection
	collection2 := query.NewMongoCollection(db,"BookTwo") // Create Another instance of Collection

	//book := model.BookTwo{"5", "1111111", "Cumilla"}

	//fmt.Println(collection.Insert(book))
	fmt.Println(collection2.DeleteById("5"))
}