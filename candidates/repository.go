package candidates

import (
	"context"
	"elections-api/database"
	"elections-api/graph/model"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCandidate(candidate model.NewCandidate) (*model.Candidate, error) {
	candidateCollection := database.MI.DB.Collection(os.Getenv("MONGO_CANDIDATES_COLLECTION"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	result, err := candidateCollection.InsertOne(ctx, &model.NewCandidateDatabase{
		Name: candidate.Name,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &model.Candidate{ 
		ID: result.InsertedID.(primitive.ObjectID).Hex(),
		Name: candidate.Name,
		CreatedAt: time.Now(),
	 }, nil
}

func GetCandidates() ([]*model.Candidate, error) {
	candidateCollection := database.MI.DB.Collection(os.Getenv("MONGO_CANDIDATES_COLLECTION"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	result, err := candidateCollection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	var candidates []*model.Candidate

	for result.Next(ctx) {
		var candidate *model.Candidate
		err := result.Decode(&candidate)
		
		if err != nil {
			log.Fatal(err)
		}

		candidates = append(candidates, candidate)
	}

	return candidates, nil
}