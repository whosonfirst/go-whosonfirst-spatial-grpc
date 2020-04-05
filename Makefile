proto:
	 protoc -I spatial/ spatial/spatial.proto --go_out=plugins=grpc:spatial

debug:
	go run -mod vendor cmd/server/main.go -mode directory:// /usr/local/data/sfomuseum-data-maps/data
