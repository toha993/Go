package main

import (
	config "Package/Generics/pkg/config"
	//"Package/Generics/pkg/model"
	query "Package/Generics/pkg/query"

	model "Package/Generics/pkg/model"
	"fmt"
)

func main() {
	db,_ := config.MongoConnection("Books") // Create instance of Database
	collection := query.NewMongoCollection(db,"book") // Create instance of Collection
	//collection2 := query.NewMongoCollection(db,"BookTwo") // Create Another instance of Collection

	//book := model.Book{book.Id:"101", book.Title:"Don'tKnow",book.Authorname:"NoName"}
	book := model.Book{
		Title: "Don'tKnow",
		Id: "102",
		Authorname: "NILLLL",
	}

	id := "61486c2daa42c647bd0a695d"
	fmt.Println(collection.Save(book,id))
}