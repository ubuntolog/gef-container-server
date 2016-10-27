SOURCES := $(shell find . -iname '*.go')

build: gef-container-server

build-linux: gef-container-server-linux

gef-container-server-linux: $(SOURCES)
	#golint ./...
	#go vet ./...
	# go test ./...
	cd ./src
# update the packages
	#GOOS=linux GOARCH=amd64 go install github.com/eudat-gef/gef-docker/dckr
	#GOOS=linux GOARCH=amd64 go install github.com/eudat-gef/gef-docker/server
# install the bin
	GOOS=linux GOARCH=amd64 go install github.com/eudat-gef/gef-container-server

gef-container-server: $(SOURCES)
	#golint ./...
	#go vet ./...
	# go test ./...
	cd ./src
# update the packages
	#go install github.com/eudat-gef/gef-docker/dckr
	#go install github.com/eudat-gef/gef-docker/server
# install the bin
	go install github.com/eudat-gef/gef-container-servercd


clean:
	go clean
