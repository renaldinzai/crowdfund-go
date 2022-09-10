package service

import (
	"crowdfund-go/campaign"
	"crowdfund-go/campaign/request"
	"crowdfund-go/entity"

	"fmt"

	"github.com/gosimple/slug"
)

type service struct {
	repository campaign.Repository
}

func NewService(repository campaign.Repository) *service {
	return &service{repository}
}

func (s *service) Campaigns(userID int) ([]entity.Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) Create(req request.Create) (entity.Campaign, error) {
	campaign := entity.Campaign{}
	campaign.Name = req.Name
	campaign.ShortDescription = req.ShortDescription
	campaign.Description = req.Description
	campaign.Perks = req.Perks
	campaign.GoalAmount = req.GoalAmount
	campaign.UserID = req.User.ID

	slugCandidate := fmt.Sprintf("%s %d", req.Name, req.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}
