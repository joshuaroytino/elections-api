package repository

import (
	"context"
	"elections-api/custom_model"
	"elections-api/database"
	"elections-api/graph/model"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateCandidate(candidate model.NewCandidateRequest) (*custom_model.Candidate, error) {
	var createdCandidate *custom_model.Candidate

	candidateCollection := database.MI.DB.Collection(os.Getenv("MONGO_CANDIDATES_COLLECTION"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cursor, err := candidateCollection.InsertOne(ctx, &custom_model.NewCandidateDTO{
		Name: candidate.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Fatal(err)
		return createdCandidate, err
	}

	candidateID := cursor.InsertedID.(primitive.ObjectID)

	candidateCollection.FindOne(ctx, bson.D{{"_id", candidateID}}).Decode(&createdCandidate)

	return createdCandidate, nil

}

func GetCandidates() ([]*custom_model.Candidate, error) {
	var candidates []*custom_model.Candidate

	candidateCollection := database.MI.DB.Collection(os.Getenv("MONGO_CANDIDATES_COLLECTION"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	result, err := candidateCollection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
		return candidates, err
	}

	for result.Next(ctx) {
		var candidate *custom_model.Candidate
		err := result.Decode(&candidate)
		
		if err != nil {
			log.Fatal(err)
			return candidates, err
		}

		candidates = append(candidates, candidate)
	}

	return candidates, nil
}

func GetCandidate(id primitive.ObjectID) (*custom_model.Candidate, error) {
	var candidate *custom_model.Candidate

	candidateCollection := database.MI.DB.Collection(os.Getenv("MONGO_CANDIDATES_COLLECTION"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	result := candidateCollection.FindOne(ctx, bson.M{"_id": id})

	err := result.Decode(&candidate)

	if err != nil {
		log.Fatal(err)
		return candidate, err
	}

	return candidate, nil
}

func UpdateCandidate(id primitive.ObjectID, input model.UpdateCandidateRequest) (*custom_model.Candidate, error) {
	var candidate *custom_model.Candidate

	candidateCollection := database.MI.DB.Collection(os.Getenv("MONGO_CANDIDATES_COLLECTION"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	result := candidateCollection.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": input}, options.FindOneAndUpdate().SetReturnDocument(1))

	err := result.Decode(&candidate)

	if err != nil {
		log.Fatal(err)
		return candidate, err
	}

	return candidate, nil
}