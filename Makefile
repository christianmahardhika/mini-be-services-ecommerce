.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/main

test:
	go test ./domain/... -v -cover

clean:
	rm -rf ./bin

docker-start:
	docker-compose up --build -d

docker-stop:
	docker-compose down