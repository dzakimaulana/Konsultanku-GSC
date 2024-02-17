build:
	go build -o ./bin/konsultanku ./cmd/konsultanku/main.go

run: build
	./bin/konsultanku