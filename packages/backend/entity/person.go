package entity

import (
	"go-tp-docker/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Person struct {
	ID primitive.ObjectID `bson:"_id, omitempty"`
	Name string `bson:"name" json:"name"`
}

func Get() (*[]Person, error) {
	client, ctx, cancel := db.GetMongoConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	var people []Person

	cursor, _ := client.Database("tp-docker").Collection("name").Find(ctx, bson.D{})

	for cursor.Next(ctx) {
		var person Person
		err := cursor.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		people = append(people, person)
	}

	return &people, nil
}

func Create(person *Person) (primitive.ObjectID, error) {
	client, ctx, cancel := db.GetMongoConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	person.ID = primitive.NewObjectID()

	result, err := client.Database("tp-docker").Collection("name").InsertOne(ctx, person)
	if err != nil {
		log.Printf("Could not create Person: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}