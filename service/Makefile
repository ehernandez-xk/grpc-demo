.PHONY: setup compile compile-go compile-java

setup:
	go get -u github.com/golang/protobuf/protoc-gen-go

compile: compile-go compile-java

compile-go:
	protoc --go_out=plugins=grpc:. service.proto

compile-java:
	#todo
