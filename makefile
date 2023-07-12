compose-up:
	echo "Starting docker environment"
	docker-compose -f docker-compose.yml up --build
.PHONY: compose-up

compose-down:
	docker-compose down --remove-orphans
.PHONY: compose-down

run:
	go run main.go
.PHONY: run
