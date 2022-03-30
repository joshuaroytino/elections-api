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