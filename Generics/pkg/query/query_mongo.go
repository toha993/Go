package query

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCollection struct{
	db *mongo.Database
	collection string
}

func NewMongoCollection(db *mongo.Database, collection string) *MongoCollection{
	return &MongoCollection{
		db: db,
		collection: collection,
	}
}

type Operation interface{
	GetAll() (interface{}, error)
	GetId(string) (interface{}, error)
	DeleteById(string) error
	Insert(interface{}) error
	Save(interface{},string) error
}

func (r *MongoCollection) Save(data interface{},Id string) (error){
	id, _ := primitive.ObjectIDFromHex(Id)
	filter := bson.M{"_id": id}
	opts := options.Update().SetUpsert(true)

	update := bson.D{primitive.E{Key: "$set", Value: data}}
	
	_, err := r.db.Collection(r.collection).UpdateOne(context.TODO(), filter, update, opts)
	return err
}

func (r *MongoCollection) Insert(data interface{}) (error){
	_, err := r.db.Collection(r.collection).InsertOne(context.TODO(), data)


	if err != nil{
		return err
	}
	return nil
}

func (r *MongoCollection) GetAll() (interface{}, error){
	cur, err := r.db.Collection(r.collection).Find(context.TODO(), bson.M{})

	if err != nil{
		log.Fatal(err)
	}

	var result []interface{}

	for cur.Next(context.TODO()){
		var elem interface{}

		cur.Decode(&elem)
		result = append(result, elem)
	}

	return result,nil
}

func (r *MongoCollection) GetId(id string) (interface{}, error){
	var result interface{}
	filter := bson.M{"id": id}
	err := r.db.Collection(r.collection).FindOne(context.TODO(), filter).Decode(&result)

	if err != nil{
		log.Fatal(err)
	}

	return result,nil
}

func (r *MongoCollection) DeleteById(id string) error{

	filter := bson.M{"id": id}
	_, err := r.db.Collection(r.collection).DeleteOne(context.TODO(), filter)
	return err
}