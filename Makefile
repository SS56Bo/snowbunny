build:
	@go build

run: build
	 @.bin/app
	 @echo "Build done"