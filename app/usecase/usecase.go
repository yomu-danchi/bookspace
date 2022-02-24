package usecase

import (
	"github.com/yuonoda/bookspace/app/domain/repositories"
	"github.com/yuonoda/bookspace/app/usecase/dto"
)

type Usecase struct {
	repository repositories.Repository
}

func (u *Usecase) CreateUser(dtoUser dto.User) (dto.User, error) {
	return dto.User{}, nil
}
