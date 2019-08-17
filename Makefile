env:
	docker-compose up -d

run:
	go run ./bin/launcher/main.go

build:
	docker build -t profile:1.0 .

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*