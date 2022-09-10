package campaign

import (
	"crowdfund-go/campaign/request"
	"crowdfund-go/entity"
)

type Service interface {
	Campaigns(userID int) ([]entity.Campaign, error)
	Create(input request.Create) (entity.Campaign, error)
}
