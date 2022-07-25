build:
	go mod vendor
	go build -mod=vendor -o main main.go

run:
	go run -mod=vendor main.go
