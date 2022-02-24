package usecase

import (
	"github.com/yuonoda/bookspace/app/domain/models/user"
	"github.com/yuonoda/bookspace/app/domain/repositories"
	"github.com/yuonoda/bookspace/app/errors"
	"github.com/yuonoda/bookspace/app/usecase/dto"
	"golang.org/x/xerrors"
)

type Usecase struct {
	repository repositories.Repository
}

func (u *Usecase) CreateUser(dtoUser dto.User) (dto.User, error) {
	if dtoUser.ID != "" {
		return dto.User{}, errors.Invalid(xerrors.Errorf("user id cannot exist, id: %s", dtoUser.ID))
	}

	newUserID, err := user.GenID()
	if err != nil {
		return dto.User{}, err
	}
	newUserName := user.NewName(dtoUser.Name)
	newUser := user.NewUser(newUserID, newUserName)

	if err := u.repository.SaveUser(newUser); err != nil {
		return dto.User{}, err
	}

	createdUser := dto.User{
		ID:   newUserID.String(),
		Name: newUserName.String(),
	}
	return createdUser, nil
}
