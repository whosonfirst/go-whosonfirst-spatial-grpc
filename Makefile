proto:
	 protoc -I spatial/ spatial/spatial.proto --go_out=plugins=grpc:spatial

cli:
	go build -mod vendor -o bin/server cmd/server/main.go
	go build -mod vendor -o bin/client cmd/client/main.go

debug:
	go run -mod vendor cmd/server/main.go -mode directory:// /usr/local/data/sfomuseum-data-maps/data
