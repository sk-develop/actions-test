build:
	docker-compose build

up:
	docker-compose up -d --remove-orphans

down:
	docker-compose down

restart:
	docker-compose restart golang && docker-compose ps

logs:
	docker-compose logs --tail=10 -f golang

test:
	docker-compose exec golang go test -cover ./...
