.PHONY: build clean clean-untagged

build:
	docker build -t trendly-backend .
	docker run -d -p 8080:8080 trendly-backend

clean:
	docker container prune -f
	docker images prune -a
	docker volume prune -f
	docker rmi $(shell docker images | awk '/^<none>/ {print $$3}')

