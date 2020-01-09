package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	
	"c:/users/g/downloads/compressed/go_mongo-master/go_mongo-master/vendor/go.mongodb.org/mongo-driver/bson"
	"c:/users/g/downloads/compressed/go_mongo-master/go_mongo-master/vendor/go.mongodb.org/mongo-driver/bson/primitive"
	"c:/users/g/downloads/compressed/go_mongo-master/go_mongo-master/vendor/go.mongodb.org/mongo-driver/mongo"
	"c:/users/g/downloads/compressed/go_mongo-master/go_mongo-master/vendor/go.mongodb.org/mongo-driver/mongo/options"
	"C:/Users/g/Downloads/Compressed/go_mongo-master/go_mongo-master/vendor/go.mongodb.org/mongo-driver/mongo/readpref"
) k

var client *mongo.Client

func dbConn() (client *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		panic(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("DB Connected sucessfully !")
	return client

}

func init() {
	client = dbConn()
}

type LandLord struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName        string             `json:"FullName,omitempty" bson:"FullName,omitempty"`
	UserName        string             `json:"UserName,omitempty" bson:"UserName,omitempty"`
	Email           string             `json:"Email,omitempty" bson:"Email,omitempty"`
	HouseNo         int                `json:"HouseNo,omitempty" bson:"HouseNo,omitempty"`
	Country         string             `json:"Country,omitempty" bson:"Country,omitempty"`
	State           string             `json:"State,omitempty" bson:"State,omitempty"`
	City            string             `json:"City,omitempty" bson:"City,omitempty"`
	Password        string             `json:"Password,omitempty" bson:"Password,omitempty"`
	ConfirmPassword string             `json:"ConfirmPassword,omitempty" bson:"ConfirmPassword,omitempty"`
}

func usersHandler(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodPost {
		response.Header().Set("content-type", "application/json")
		var people []LandLord
		collection := client.Database("class").Collection("user")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var LandLord LandLord
			cursor.Decode(&LandLord)
			people = append(people, LandLord)
		}
		if err := cursor.Err(); err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(response).Encode(people)
	} else if request.Method == http.MethodPost {
		response.Header().Set("content-type", "application/json")
		var LandLord LandLord
		_ = json.NewDecoder(request.Body).Decode(&LandLord)

		collection := client.Database("class").Collection("user")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		result, _ := collection.InsertOne(ctx, LandLord)
		json.NewEncoder(response).Encode(result)
	} else {
		response.Write([]byte(`{ "message": "` + request.Method + `" }`))
	}

}

func userHandler(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodPost {
		response.Header().Set("content-type", "application/json")
		value := request.URL.Query().Get("id")
		id, _ := primitive.ObjectIDFromHex(value)
		var LandLord LandLord
		collection := client.Database("class").Collection("user")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := collection.FindOne(ctx, LandLord{ID: id}).Decode(&LandLord)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(response).Encode(LandLord)
	} else if request.Method == http.MethodPatch {
		response.Header().Set("content-type", "application/json")
		value := request.URL.Query().Get("id")
		id, _ := primitive.ObjectIDFromHex(value)

		collection := client.Database("class").Collection("user")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

		var LandLord LandLord
		_ = json.NewDecoder(request.Body).Decode(&LandLord)
		update := bson.M{"$set": LandLord}

		result, err := collection.UpdateOne(ctx, LandLord{ID: id}, update)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(response).Encode(result)
	} else if request.Method == http.MethodDelete {
		response.Header().Set("content-type", "application/json")
		value := request.URL.Query().Get("id")
		id, _ := primitive.ObjectIDFromHex(value)

		collection := client.Database("class").Collection("user")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		result, err := collection.DeleteOne(ctx, LandLord{ID: id})
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
			return
		}
		json.NewEncoder(response).Encode(result)
	} else {
		response.Write([]byte(`{ "message": "` + request.Method + `" }`))
	}

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/user", userHandler)

	http.ListenAndServe(":8080", mux)

}
