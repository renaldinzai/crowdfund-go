package user

import (
	"crowdfund-go/entity"
	"crowdfund-go/user/request"
)

type Service interface {
	Register(request request.Register) (entity.User, error)
	Login(request request.Login) (entity.User, error)
	IsEmailAvailable(request request.CheckEmail) (bool, error)
	SaveAvatar(ID int, fileLocation string) (entity.User, error)
	GetUserByID(ID int) (entity.User, error)
}
