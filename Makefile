SHELL := /bin/bash

SRC = ./src
OUTPUT = ../../build

.phony all:
all: build

# Build Commands
.phony build:
build: build-web build-transaction build-data build-audit build-generator build-quote-mock

.phony build-web:
build-web: 
	cd $(SRC)/web && go build -o $(OUTPUT)/web

.phony build-transaction:
build-transaction: 
	cd $(SRC)/transaction && go build -o $(OUTPUT)/transaction

.phony build-data:
build-data: 
	cd $(SRC)/data && go build -o $(OUTPUT)/data

.phony build-audit:
build-audit: 
	cd $(SRC)/audit && go build -o $(OUTPUT)/audit

.phony build-generator:
build-generator: 
	cd $(SRC)/generator && go build -o $(OUTPUT)/generator

.phony build-quote-mock:
build-quote-mock: 
	cd $(SRC)/quote-mock && go build -o $(OUTPUT)/quote-mock

# Docker Compose Commands
.phony docker-deploy-dev:
docker-deploy-dev:
	docker-compose -f docker-compose.yml -f docker-compose.local.yml -f docker-compose.dev.yml up --build

.phony docker-deploy-local:
docker-deploy-local:
	docker-compose -f docker-compose.yml -f docker-compose.local.yml up --build

.phony docker-deploy-lab:
docker-deploy-lab: build
	sudo docker-compose -f docker-compose.yml -f docker-compose.lab.yml up --build

.phony docker-redeploy-dev:
docker-redeploy:
	docker-compose up -f docker-compose.yml -f docker-compose.local.yml -f docker-compose.dev.yml --build --force-recreate --no-deps -d $(c)

.phony docker-teardown:
docker-teardown:  
	docker-compose down --remove-orphans

# Docker Container Commands
.phony docker-list:
docker-list: 
	docker ps

# Container Specific Commmands
# Example: make c=CONTAINER_NAME docker-shell 
.phony docker-shell:
docker-shell:
	docker exec -it $(c) bash

.phony docker-stop:
docker-stop:
	docker stop $(c)

.phony docker-start:
docker-start:
	docker start $(c)

.phony docker-remove:
docker-remove:
	docker rm $(c)