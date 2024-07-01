GO=$(shell which go)

# run tests
test:
	go clean -testcache
	go test ./...

# generate coverage and open in html
coverage:
	go test ./... -coverprofile=c.out
	go tool cover -html="c.out"
	rm -rf c.out
