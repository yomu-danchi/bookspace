package dto

import "github.com/yuonoda/bookspace/app/domain/models/user"

type Users []User

func ToDtoUsers(domainUsers user.Users) Users {
	dtoUsers := make(Users, 0, len(domainUsers))
	for _, domainUser := range domainUsers {
		dtoUsers = append(dtoUsers, ToDtoUser(domainUser))
	}
	return dtoUsers
}
