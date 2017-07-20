.PHONY: build build-image build-lx push run

build: build-lx build-image

build-lx:
	GOOS="linux" GOARCH="amd64" go build -o goserver .

build-image:
	docker build -t registry.xumak.gt:5000/ehernandez/go-server .

push:
	docker push registry.xumak.gt:5000/ehernandez/go-server

clean:
	#Cleaning the local docker stuff
	docker stop go-server || true
	docker rmi registry.xumak.gt:5000/ehernandez/go-server || true

run:
	docker run -d --rm --name go-server -p 50051:50051 registry.xumak.gt:5000/ehernandez/go-server
