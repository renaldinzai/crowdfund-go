package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crowdfund-go/auth"
	"crowdfund-go/campaign"
	"crowdfund-go/graph/generated"
	"crowdfund-go/graph/model"
	"crowdfund-go/user"
	"fmt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	repository := user.NewRepository(r.DB)
	userService := user.NewService(repository)
	authService := auth.NewService()

	inputFormat := user.RegisterUserInput{
		Name:       input.Name,
		Occupation: input.Occupation,
		Email:      input.Email,
		Password:   input.Password,
	}

	newUser, err := userService.RegisterUser(inputFormat)
	if err != nil {
		return nil, err
	}

	token, err := authService.GenerateToken(newUser.ID)
	if err != nil {
		return nil, err
	}

	result := &model.User{
		ID:         newUser.ID,
		Name:       newUser.Name,
		Occupation: newUser.Occupation,
		Email:      newUser.Email,
		Token:      token,
	}

	return result, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	repository := user.NewRepository(r.DB)
	userService := user.NewService(repository)
	authService := auth.NewService()

	inputFormat := user.LoginInput{
		Email:    input.Email,
		Password: input.Password,
	}

	loggedinUser, err := userService.Login(inputFormat)
	if err != nil {
		return "", err
	}

	token, err := authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// CreateCampaign is the resolver for the createCampaign field.
func (r *mutationResolver) CreateCampaign(ctx context.Context, input model.NewCampaign) (*model.Campaign, error) {
	repository := campaign.NewRepository(r.DB)
	service := campaign.NewService(repository)

	currentUser := auth.ForContext(ctx)
	if currentUser == (user.User{}) {
		return nil, fmt.Errorf("access denied")
	}

	inputFormat := campaign.CreateCampaignInput{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		GoalAmount:       input.GoalAmount,
		Perks:            input.Perks,
		User:             currentUser,
	}

	newCampaign, err := service.CreateCampaign(inputFormat)
	if err != nil {
		return nil, err
	}

	result := &model.Campaign{
		ID:               newCampaign.ID,
		UserID:           newCampaign.UserID,
		Name:             newCampaign.Name,
		ShortDescription: newCampaign.ShortDescription,
		GoalAmount:       newCampaign.GoalAmount,
		Slug:             newCampaign.Slug,
	}

	return result, nil
}

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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
