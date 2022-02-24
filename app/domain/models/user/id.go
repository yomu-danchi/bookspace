package user

import gonanoid "github.com/matoous/go-nanoid/v2"

type ID string

func NewID(ui string) ID {
	return ID(ui)
}

func GenID() (ID, error) {
	id, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	return NewID(id), nil
}

func (id ID) String() string {
	return string(id)
}
