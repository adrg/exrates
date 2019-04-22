all: install
	go test .
	go vet .
	@echo "go fmt ." && test -z $(shell gofmt -d -e -l . | tee /dev/stderr)
	find . -name '*.go' | xargs ineffassign
	staticcheck .
	unconvert -v .
	unparam .
	misspell -locale GB -error *.md *.go
	golint .

install:
	go get github.com/gordonklaus/ineffassign
	go get honnef.co/go/tools/cmd/staticcheck
	go get github.com/mdempsky/unconvert
	go get mvdan.cc/unparam
	go get github.com/client9/misspell/cmd/misspell
	go get golang.org/x/lint/golint
