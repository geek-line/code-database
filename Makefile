build:
	@echo "build for production"
	GOARCH="amd64" GOOS="linux" go build

start: build
	@echo "start backend server for production!"
	BUILD_MODE=prod ./code-database

dev:
	@echo "start backend server for development!"
	BUILD_MODE=dev air