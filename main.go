package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoURI          string
	DatabaseName      string
	CollectionNames   []string
	ChangeStreamStage bson.D
}

func main() {
	CreateNewFlow()

	// Set up configuration
	config := Config{
		MongoURI:        "<db_uri>",
		DatabaseName:    "go-watch-stream",
		CollectionNames: []string{"users", "products"},
		ChangeStreamStage: bson.D{
			{"$match", bson.D{
				{"operationType", bson.D{
					{"$in", bson.A{"insert", "update", "delete", "replace"}},
				}},
			}},
		},
	}

	// Set up MongoDB client and connect to the database
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		panic(err)
	}
	// Connect to MongoDB
	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	// Open a change stream on the specified collections
	for _, collectionName := range config.CollectionNames {
		collection := client.Database(config.DatabaseName).Collection(collectionName)
		go watchCollection(collection, config.ChangeStreamStage)
	}

	// Wait for a signal to exit
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs
	log.Println("Exiting...")
}

func watchCollection(collection *mongo.Collection, changeStreamStage bson.D) {
	// Set up options for the change stream
	pipeline := mongo.Pipeline{changeStreamStage}

	changeStream, err := collection.Watch(context.Background(), pipeline, nil)
	if err != nil {
		panic(err)
	}
	defer changeStream.Close(context.Background())

	for changeStream.Next(context.Background()) {
		var changeDocument bson.M
		if err = changeStream.Decode(&changeDocument); err != nil {
			panic(err)
		}

		parseChangeDocument(changeDocument)
	}
}

func parseChangeDocument(changeDocument bson.M) {
	changeBytes, err := json.Marshal(changeDocument)
	if err != nil {
		log.Println(err)
	}

	var parsedDocument map[string]interface{}
	if err = json.Unmarshal(changeBytes, &parsedDocument); err != nil {
		log.Println(err)
	}

	operationType := parsedDocument["operationType"]
	data := parsedDocument["fullDocument"]
	collection := parsedDocument["ns"].(map[string]interface{})["coll"].(string)

	RunFlow(collection, operationType, data)
}
