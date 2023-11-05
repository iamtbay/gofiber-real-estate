run: build
	./realestate

build:
	@go build -o realestate ./cmd/*.go