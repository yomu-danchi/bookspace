package firestore

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

const (
	BooksCollectionName = "books"
	UsersCollectionName = "users"
)
