package book

type Title string

func NewTitle(s string) Title {
	return Title(s)
}

func (n Title) String() string {
	return string(n)
}
