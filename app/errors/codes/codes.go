package codes

type Code int

const (
	OK Code = iota
	Invalid
	BadParams
	InternalError
)
