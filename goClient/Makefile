.PHONY: build build-lx build-image push clean run-add run-list run-fake

build: build-lx build-image

build-lx:
	GOOS="linux" GOARCH="amd64" go build -o goclient .

build-image:
	docker build -t registry.xumak.gt:5000/ehernandez/go-client .

push:
	docker push registry.xumak.gt:5000/ehernandez/go-client

clean:
	#Cleaning the local docker stuff
	docker rmi registry.xumak.gt:5000/ehernandez/go-client || true

run-add:
	docker run --rm registry.xumak.gt:5000/ehernandez/go-client -option add -target 10.4.4.118 -name "Eddy Hernandez"

run-list:
	docker run --rm registry.xumak.gt:5000/ehernandez/go-client -target 10.4.4.118

run-fake:
	docker run --rm registry.xumak.gt:5000/ehernandez/go-client -option fake -target 10.4.4.118
