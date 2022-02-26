package firestore

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

const (
	booksCollectionName = "books"
	usersCollectionName = "users"
)
