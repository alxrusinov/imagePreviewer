up:	docker_build docker_run

run:
	go run 'cmd/app/main.go'

docker_build:
	docker build . -t image-previewer

docker_run:
	docker run -p 80:3000 --name image-previewer image-previewer:latest

docker_container_clean:
	docker rm image-previewer