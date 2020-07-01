cli:
	go build -mod vendor -o bin/oembed-populate cmd/oembed-populate/main.go
	go build -mod vendor -o bin/oembed-server cmd/oembed-server/main.go
