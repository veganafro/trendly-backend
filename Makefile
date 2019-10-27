.PHONY: build run clean

build:
	docker build --rm -t trendly-backend .

run:
	docker run -d --name trendly-backend -p 8080:8080 trendly-backend

clean:
	docker container prune -f
	docker images prune -a
	docker volume prune -f
	docker rmi $(shell docker images | awk '/^<none>/ {print $$3}')

