package user

import "crowdfund-go/entity"

type Repository interface {
	Save(user entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
	FindByID(ID int) (entity.User, error)
	Update(user entity.User) (entity.User, error)
}
