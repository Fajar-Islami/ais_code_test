enterpostgre:
	docker exec -it postgres_ais psql -U admin ais_db

enterredis:
	docker exec -it redis_ais redis-cli

compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app

composerestart:
	docker-compose down -v
	docker-compose up -d

dockerrun:
	docker build . -t ais_be:1.0.0
	docker run --rm --name ais_be ais_be:1.0.0

dockerclear:
	docker stop ais_be
	docker rm ais_be
	docker rmi ais_be:1.0.0

# dockerrestart:
# 	make dockerrun
# go test -v ./app