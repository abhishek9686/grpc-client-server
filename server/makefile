.PHONY: all watch build-linux build-os dockerize clean

all: 
	go mod tidy
	${MAKE} build-os dockerize clean
build-linux:
	go mod tidy
	GOOS=linux CGO_ENABLED=0 go build -o usersServer.linux
build-os:
	go mod tidy
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=0 go build  -o usersServer.exe
else
	CGO_ENABLED=0 go build  -o usersServer
endif

dockerize:
	${MAKE} build-linux
	docker build -f ../build/Dockerfile.server . -t abhi9686/users:v1

watch:
	reflex -s -r '\.go$$' make run

clean:
	go clean
ifeq ($(OS),Windows_NT)
	@rm -f ./usersServer
else
	@rm -f .\usersServer.exe
endif