.PHONY: build push clean run-add run-list

build:
	docker build -t registry.xumak.gt:5000/ehernandez/node-client .

push:
	docker push registry.xumak.gt:5000/ehernandez/node-client

clean:
	#Cleaning the local docker stuff
	docker rmi registry.xumak.gt:5000/ehernandez/node-client || true

run-add:
	docker run --rm registry.xumak.gt:5000/ehernandez/node-client --option add --target "10.4.4.118" --name "Luis Morales"

run-list:
	docker run --rm registry.xumak.gt:5000/ehernandez/node-client --target "10.4.4.118"
