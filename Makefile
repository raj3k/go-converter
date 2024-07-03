docker-compose: docker-compose-stop docker-compose-start
.PHONY: docker-compose

docker-compose-start:
	docker-compose -f docker-compose-base.yaml -f docker-compose.yaml up --build
.PHONY: docker-compose-start

docker-compose-stop:
	docker-compose -f docker-compose-base.yaml -f docker-compose.yaml down --remove-orphans -v
.PHONY: docker-compose-stop

docker-compose-base: docker-compose-base-stop docker-compose-base-start
.PHONY: docker-compose-base

docker-compose-base-start:
	docker-compose -f docker-compose-base.yaml up -d --build
.PHONY: docker-compose-base-start

docker-compose-base-stop:
	docker-compose -f docker-compose-base.yaml down --remove-orphans -v