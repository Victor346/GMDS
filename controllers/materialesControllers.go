package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strconv"
	"testApi/models"
)

type MaterialDB struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Nombre      string
	Descripcion string
	Codigo      string
}

var GetOneMaterial = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idMaterial := params["ID"]

	client := models.GetClient
	materialesCollection := client().Database("GMDS").Collection("Materiales")

	idObjectId, err := primitive.ObjectIDFromHex(idMaterial)

	if err != nil {
		//Todo write correct status code
		w.WriteHeader(404)
		return
	}

	var result MaterialDB
	filter := bson.D{{"_id", idObjectId}}

	err = materialesCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	js, err := json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	fmt.Printf("Found a single document: %+v\n", result)
}

var GetMultipleMaterials = func(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	lowLimitStr := params["LowLimit"]
	upperLimitStr := params["UpLimit"]

	lowLimit, err := strconv.Atoi(lowLimitStr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	upperLimit, err := strconv.Atoi(upperLimitStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	numOfDocuments := upperLimit - lowLimit

	if numOfDocuments <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	client := models.GetClient
	materialesCollection := client().Database("GMDS").Collection("Materiales")

	findOptions := options.Find()
	findOptions.SetLimit(int64(numOfDocuments))
	findOptions.SetSkip(int64(lowLimit))

	var results []*MaterialDB

	cur, err := materialesCollection.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for cur.Next(context.TODO()) {
		 var elem MaterialDB
		 err := cur.Decode(&elem)
		 if err != nil {
		 	fmt.Println("Hola")
			 continue
		 }

		 results = append(results, &elem)
	}

	fmt.Println(results)
	fmt.Println(lowLimit)
	fmt.Println(upperLimit)
}
