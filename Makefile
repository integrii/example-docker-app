build:
	docker build -t integrii/example-docker-app .

push:
	docker push integrii/example-docker-app

run:
	docker run -p 8085:80 --rm integrii/example-docker-app
