run:
	@DSN=your-dsn-to-mongo-db go run .

build:
	CGO_ENABLED=0 GOOS=linux go build .