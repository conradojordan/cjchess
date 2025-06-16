help:
	@echo 'Available make commands:'
	@echo '    help (default)     - prints this help message'
	@echo '    build              - compiles the project into a binary'
	@echo '    run                - compiles the project into a temporary binary file and run it'
	@echo '    test               - run all tests'

build:
	go build ./... -o cjchess

run:
	go run .

test:
	go test -v ./...
