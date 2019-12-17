get_deps:
	echo ### Downloading dependencies ###
	go mod download

test: get_deps
	echo ### Running unit tests ###
	cd pkg/actions
	go test ./...

vet: get_deps
	go vet ./...

prereqs: test

prep: vet prereqs



