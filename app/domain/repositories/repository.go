package repositories

import (
	"github.com/yuonoda/bookspace/app/domain/models/user"
)

type Repository interface {
	SaveUser(user user.User) error
}
