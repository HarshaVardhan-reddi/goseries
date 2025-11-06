package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongodbtutorial/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.5.9"
const dbName = "Netflix"
const colName = "watchlist"

var collection *mongo.Collection

// connect with mongo
// this function will be executed only once after the code execution
func init(){
	// cleint conn
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(clientOptions)
	if err != nil{
		log.Fatal(err)
	}
	
	if err := client.Ping(context.TODO(), clientOptions.ReadPreference); err != nil {
		panic(err)
	}
	fmt.Println("Mongo connection established successfully")
	collection = client.Database(dbName).Collection(colName)
	// if collection is ready
	fmt.Println("Collection instance is ready")
}

// Mongo db helpers
func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Movie",movie)
	fmt.Println("Inserted the record with id", inserted.InsertedID)
}

func updateOneMovie(movieId string){
	id, err := primitive.ObjectIDFromHex(movieId)
	if(err != nil){
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set":bson.M{"watched":true}}

	result, errupdating := collection.UpdateOne(context.Background(), filter, update)
	if errupdating != nil {
		log.Fatal(errupdating)
	}
	fmt.Println("updated the movie with id and other details:", result.Acknowledged, result.UpsertedID, result.ModifiedCount)
}


func deletionOneMovie(movieId string){
	id, err := bson.ObjectIDFromHex(movieId)
	if err != nil{
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	result, errordeletion := collection.DeleteOne(context.Background(),filter)
	if errordeletion != nil{
		log.Fatal(errordeletion)
	}
	fmt.Println("deleted the following", result.DeletedCount)
}


func deleteAllTheMovies() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("The result here is:", result.DeletedCount)
	return result.DeletedCount
}


func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.TODO(),bson.M{})
	if err != nil{
		log.Fatal(err)
	}
	if cursorerr := cursor.Err(); cursorerr != nil{
		log.Fatal(cursorerr)
	}
	fmt.Println("Cursor object got returned", cursor)
	var movies []primitive.M
	for cursor.Next(context.TODO()){
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil{
			log.Fatal(err)
		}
		movies = append(movies, primitive.M(movie))
	}
	defer cursor.Close(context.Background())
	return movies
}

func GetAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)
}

func CreateAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	if !movie.IsValid(){
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message":"Invalid request missing params"})
	}
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkMovieAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if params["Id"] == ""{
		json.NewEncoder(w).Encode(map[string]string{"message":"Id is missing in the params"})
	}
	updateOneMovie(params["Id"])
	json.NewEncoder(w).Encode(map[string]string{"Update movie with id":params["Id"]})
}


func DeleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if params["Id"] == ""{
		json.NewEncoder(w).Encode(map[string]string{"message":"Id is missing in the params"})
	}
	deletionOneMovie(params["Id"])
	json.NewEncoder(w).Encode(map[string]string{"message":"deleted movie successfully"})
}


func DeleteAllMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	count := deleteAllTheMovies()
	fmt.Println("Deleting the count of movies",count)

	json.NewEncoder(w).Encode(map[string]string{"message":"deleted movie successfully"})
}