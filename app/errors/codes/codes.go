package codes

type Code int

const (
	OK Code = iota
	NotFound
	Invalid
	BadParams
	InternalError
)
