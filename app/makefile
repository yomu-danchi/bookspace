mockgen:
	moq -out ./domain/repositories/mock/mock.go -pkg mock ./domain/repositories Repository
test:
	go test ./... --count=1

init:
	go run ./scripts/init
