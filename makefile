all: 
	go mod tidy
	${MAKE} build-dockerize-server build-client
build-dockerize-server:
	${MAKE} -C server all
build-server:
	${MAKE} -C server build-os
build-client:
	${MAKE} -C client build-os
 	