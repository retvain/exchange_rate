include .env
export $(shell grep -E -v '^\s*#' .env | xargs)

up:
	docker-compose up -d

down:
	docker-compose down