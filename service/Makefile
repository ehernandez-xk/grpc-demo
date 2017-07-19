.PHONY: compile compile-go compile-java

compile: compile-go compile-java

compile-go:
	protoc --go_out=plugins=grpc:. service.proto

compile-java:
	#todo
