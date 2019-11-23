all: clean

clean: 
	find . -name 'menu' -type f -exec rm -f {} \; 
	go clean

go-get:
	rm -rf src/github.com
	rm -rf src/gopkg.in
	go get -v github.com/codegangsta/negroni
	go get -v github.com/streadway/amqp
	go get -v github.com/gorilla/mux
	go get -v github.com/unrolled/render
	go get -v github.com/satori/go.uuid
	go get -v gopkg.in/mgo.v2
	go get -v gopkg.in/mgo.v2/bson

run:
	go run src/app/$(app).go

main:
	go run src/app/main.go

format:
	go fmt menu

install:
	go install menu

build:
	go build menu

start:
	./gumball 

test-ping:
	curl localhost:8001/menu/ping

test-menu:
	curl localhost:8001/menu

docker-build: 
	docker build -t menu .
	docker images

network-create:
	docker network create --driver bridge menu

network-inspect:
	docker network inspect menu

rabbitmq-run:
	docker run --name rabbitmq \
	           --network gumball \
			   -p 8080:15672 -p 4369:4369 -p 5672:5672 \
			   -d rabbitmq:3-management
mongodb-run:
	docker run --name mongodb --network menu -p 27017:27017 -d mongo

docker-run:
	docker run \
            --network menu \
			--name gumball -p 8001:8001 -td gumball --restart always
	docker ps

docker-network:
	docker network ls

docker-network-inspect:
	docker network inspect host

docker-shell:
	docker exec -it menu bash 

docker-clean:
	docker stop mongodb
	docker stop rabbitmq
	docker rm mongodb
	docker rm rabbitmq
	docker stop menu
	docker rm menu
	docker rmi menu

docker-ip:
	docker-machine ip

docker-ps:
	 docker ps --all --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}\t"

docker-ps-ports:
	 docker ps --all --format "table {{.Names}}\t{{.Ports}}\t"
