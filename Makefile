BINARY=user-service
.DEFAULT_GOAL := run

test: 
	go test -v -cover -covermode=atomic ./...

dev:
	rm -rf my.db
	redis-cli flushall
	go build -ldflags "-X main.devenv=development"
	./${BINARY}

mobile:
	rm -rf my.db
	redis-cli flushall
	go build -ldflags "-X main.devenv=mobile"
	./${BINARY}

prod:
	go build -ldflags "-X main.devenv=production"
	./${BINARY};

docker:
	docker build -f Dockerfile -t server:lastest 
	docker run --name server -p 8080:8080 server:lastest

docker_env: 
	docker run --name server -p 8080:8080 \
	-e SERVER_HOST=127.0.0.1 \
	-e SERVER_PORT=8080 \
	-e SERVER_READ_TIMEOUT=30 \
	-e SERVER_READ_HEADER_TIMEOUT=15 \
	-e SERVER_WRITE_TIMEOUT=10 \
	-e SERVER_IDLE_TIMEOUT=10 \
	-e SERVER_MAX_HEADER_BYTES=255 \
	-e DATABASE_TYPE=mysql \
	-e DATABASE_USER=root \
	-e DATABASE_PASSWORD=ducthang \
	-e DATABASE_HOST=127.0.0.1:3306 \
	-e DATABASE_NAME=ieltscenter \
	-e COOKIE_DOMAIN=localhost \
	-e COOKIE_HTTP_ONLY=false \
	-e COOKIE_SECURE=false \
	server:lastest

build:
	go build -o ${BINARY} *.go
	./${BINARY};

dockerc:
	docker-compose up -d --force-recreate

svm: 
	docker-compose -f docker-compose.virtual.sql.yml up -d
	docker exec -it virtual-db mysql -u root -p

.PHONY: test run docker docker_env dockerc