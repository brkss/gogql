package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/brkss/gogql/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GetAwarenesses is the resolver for the GetAwarenesses field.
func (r *queryResolver) GetAwarenesses(ctx context.Context) ([]*model.Awareness, error) {
	list, err := r.Store.GetAwarnesses(ctx)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: "Something went wrong !",
		}
	}
	var results []*model.Awareness
	for _, item := range list {
		results = append(results, &model.Awareness{
			ID:    item.ID,
			Title: item.Title,
		})
	}
	return results, nil
	//panic(fmt.Errorf("not implemented: GetAwarenesses - GetAwarenesses"))
}
