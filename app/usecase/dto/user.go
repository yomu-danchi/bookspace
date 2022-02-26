package dto

import "github.com/yuonoda/bookspace/app/domain/models/user"

type User struct {
	ID   string
	Name string
}

func ToDtoUser(domainUser user.User) User {
	return User{
		ID:   domainUser.ID.String(),
		Name: domainUser.Name.String(),
	}
}
