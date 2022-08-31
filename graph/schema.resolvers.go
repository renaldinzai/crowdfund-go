package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crowdfund-go/campaign"
	"crowdfund-go/graph/generated"
	"crowdfund-go/graph/model"
)

// Campaigns is the resolver for the campaigns field.
func (r *queryResolver) Campaigns(ctx context.Context, userID *int) ([]*model.Campaign, error) {
	var result []*model.Campaign
	repository := campaign.NewRepository(r.DB)
	service := campaign.NewService(repository)

	campaigns, err := service.GetCampaigns(*userID)
	if err != nil {
		return result, err
	}

	for _, c := range campaigns {
		result = append(result, &model.Campaign{
			ID:               c.ID,
			UserID:           c.UserID,
			Name:             c.Name,
			ShortDescription: c.ShortDescription,
			GoalAmount:       c.GoalAmount,
			CurrentAmount:    c.CurrentAmount,
			Slug:             c.Slug,
		})
	}

	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
