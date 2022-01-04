bup:
	docker-compose build
	docker-compose up

bld:
	docker-compose build

up:
	docker-compose up

run:
	swagger generate spec -o ./swagger.yaml --scan-models
	go run main.go
