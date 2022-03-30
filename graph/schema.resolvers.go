package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"elections-api/custom_model"
	"elections-api/graph/generated"
	"elections-api/graph/model"
	"elections-api/repository"
)

func (r *mutationResolver) CreateCandidate(ctx context.Context, input model.NewCandidateRequest) (*custom_model.Candidate, error) {
	return repository.CreateCandidate(input)
}

func (r *queryResolver) Candidates(ctx context.Context) ([]*custom_model.Candidate, error) {
	return repository.GetCandidates()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
