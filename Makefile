BUILD_DIR="./.bin/cli"

build:
	@echo "Building..."
	@go build -o .bin/cli cmd/cli/main.go

run: build
	@${BUILD_DIR}

test-rerun:
	@go run gotest.tools/gotestsum@latest --debug --format testname --rerun-fails --packages="./..." -- -count=2

test:
	@go run gotest.tools/gotestsum@latest --debug --format testname --debug --packages="./..." -- -count=1

test-watch:
	@go run gotest.tools/gotestsum@latest --debug --format testname --watch --packages="./..." -- -count=1

test-only:
	@go run gotest.tools/gotestsum@latest --debug --format standard-verbose --watch --packages="./..." -- --run $(TEST_NAME) -count=1
