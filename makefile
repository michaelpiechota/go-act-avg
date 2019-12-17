get_deps:
	echo ### Downloading dependencies ###
	go mod download

test: get_deps
	echo ### Running unit tests ###
	cd pkg/actions
	go test ./...

fmt:
	go fmt ./... -v

vet: get_deps
	go vet ./...

prereqs: get_deps test

prep: vet prereqs



