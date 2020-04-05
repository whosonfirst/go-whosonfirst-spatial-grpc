proto:
	 protoc -I spatial/ spatial/spatial.proto --go_out=plugins=grpc:spatial
