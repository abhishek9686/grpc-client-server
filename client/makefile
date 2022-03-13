.PHONY: all watch build-linux build-os dockerize clean

all: 
	go mod tidy
	${MAKE} build-os
build-os:
	go mod tidy
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=0 go build  -o usersClient.exe
else
	CGO_ENABLED=0 go build  -o usersClient
endif

watch:
	reflex -s -r '\.go$$' make run

clean:
	go clean
ifeq ($(OS),Windows_NT)
	@rm -f ./usersClient
else
	@rm -f .\usersClient.exe
endif