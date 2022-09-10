package response

import "crowdfund-go/entity"

type Campaign struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func Format(campaign entity.Campaign) Campaign {
	campaignFormatted := Campaign{}
	campaignFormatted.ID = campaign.ID
	campaignFormatted.UserID = campaign.UserID
	campaignFormatted.Name = campaign.Name
	campaignFormatted.ShortDescription = campaign.ShortDescription
	campaignFormatted.GoalAmount = campaign.GoalAmount
	campaignFormatted.CurrentAmount = campaign.CurrentAmount
	campaignFormatted.Slug = campaign.Slug
	campaignFormatted.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatted.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatted
}

func FormatMany(campaigns []entity.Campaign) []Campaign {
	campaignsFormatted := []Campaign{}

	for _, campaign := range campaigns {
		campaignFormatted := Format(campaign)
		campaignsFormatted = append(campaignsFormatted, campaignFormatted)
	}

	return campaignsFormatted
}
