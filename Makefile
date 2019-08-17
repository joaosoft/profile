env:
	docker-compose up -d

run:
	go run ./bin/launcher/main.go

build:
	docker build -t profile:1.0 .

push:
	docker login --username joaosoft
	docker tag profile:1.0 joaosoft/profile
	docker push joaosoft/profile

fmt:
	go fmt ./...

vet:
	go vet ./*

gometalinter:
	gometalinter ./*