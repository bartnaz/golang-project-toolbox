install:
	go build ./golang-project-toolbox/...

test:
	go test ./golang-project-toolbox/...

mod-refresh:
	go clean --testcache && \
	rm go.sum && \
	go clean --modcache && \
	go mod tidy