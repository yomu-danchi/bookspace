package user

type User struct {
	ID   ID
	Name Name
}

func NewUser(ID ID, name Name) User {
	return User{
		ID:   ID,
		Name: name,
	}
}
