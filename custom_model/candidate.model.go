package custom_model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Candidate struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
type NewCandidateDTO struct {
	Name      string    `json:"name" bson:"name"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type UpdateCandidateDTO struct {
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}