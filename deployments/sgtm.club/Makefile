.PHONY: all
all: up logs

.PHONY: up
up:
	docker-compose up -d

.PHONY: logs
logs:
	while true; do make logs-once; sleep 1; done

.PHONY: logs-once
logs-once:
	docker-compose logs --tail=1000 -f

.PHONY: ps pu