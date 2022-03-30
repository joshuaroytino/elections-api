package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"elections-api/custom_model"
	"elections-api/graph/generated"
	"elections-api/graph/model"
	"elections-api/repository"
	"fmt"
	"time"
)

func (r *candidateResolver) UpdatedAt(ctx context.Context, obj *custom_model.Candidate) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCandidate(ctx context.Context, input model.NewCandidate) (*custom_model.Candidate, error) {
	return repository.CreateCandidate(input)
}

func (r *queryResolver) Candidates(ctx context.Context) ([]*custom_model.Candidate, error) {
	return repository.GetCandidates()
}

// Candidate returns generated.CandidateResolver implementation.
func (r *Resolver) Candidate() generated.CandidateResolver { return &candidateResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type candidateResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
