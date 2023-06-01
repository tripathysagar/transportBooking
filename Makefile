build:
	@go build -o bin/api 

run: build
	@./bin/api

startdb:
	@docker run --name mongodb -p 27017:27017 -d docker.io/mongo:latest 

restartdb:
	@docker start mongodb

rmdb:
	@docker stop mongodb
	@docker rm mongodb