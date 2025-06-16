help:
	@echo 'Available make commands:'
	@echo '    help (default)     - prints this help message'
	@echo '    build              - compiles the project into a binary'
	@echo '    run                - compiles and runs the project (built binary is not kept)'
	@echo '    test               - runs all tests'

build:
	go build ./... -o cjchess

run:
	go run .

test:
	go test -v ./...
