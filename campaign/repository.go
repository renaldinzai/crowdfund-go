package campaign

import "crowdfund-go/entity"

type Repository interface {
	FindAll() ([]entity.Campaign, error)
	FindByUserID(userID int) ([]entity.Campaign, error)
	Save(campaign entity.Campaign) (entity.Campaign, error)
}
