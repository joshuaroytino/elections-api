package candidates

import (
	"context"
	"elections-api/database"
	"elections-api/graph/model"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCandidate(candidate model.NewCandidate) (*model.Candidate, error) {
	candidateCollection := database.MI.DB.Collection(os.Getenv("MONGO_CANDIDATES_COLLECTION"))

	result, err := candidateCollection.InsertOne(context.TODO(), &model.NewCandidateDatabase{
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