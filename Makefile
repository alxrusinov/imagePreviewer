up:	docker_build docker_run

run:
	go run 'cmd/app/main.go'

docker_build:
	docker build . -t image-previewer

docker_run:
	docker run -p 80:80 --name image-previewer image-previewer:latest

docker_container_clean:
	docker rm image-previewer

nginx_build:
	docker build . -f Dockerfile.nginx -t nginx

nginx_run:
	docker run -p 80:80  --name nginx nginx:latest

nginx_stop:
	docker stop nginx
	docker rm nginx
	docker rmi nginx