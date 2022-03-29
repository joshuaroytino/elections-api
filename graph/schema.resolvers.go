package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"elections-api/candidates"
	"elections-api/graph/generated"
	"elections-api/graph/model"
)

func (r *mutationResolver) CreateCandidate(ctx context.Context, input model.NewCandidate) (*model.Candidate, error) {
	return candidates.CreateCandidate(input)
}

func (r *queryResolver) Candidates(ctx context.Context) ([]*model.Candidate, error) {
	return candidates.GetCandidates()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
