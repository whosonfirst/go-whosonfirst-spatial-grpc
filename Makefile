# https://developers.google.com/protocol-buffers/docs/reference/go-generated
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative spatial/spatial.proto

cli:
	go build -mod vendor -o bin/server cmd/server/main.go
	go build -mod vendor -o bin/client cmd/client/main.go

debug:
	go run -mod vendor cmd/server/main.go -mode directory:// /usr/local/data/sfomuseum-data-maps/data
