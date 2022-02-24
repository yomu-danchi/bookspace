package book

type ISBN13 string

func NewISBN13(s string) ISBN13 {
	return ISBN13(s)
}

func (n ISBN13) String() string {
	return string(n)
}
