install:
	make install_protoc_linux_x86_64
install_protoc_linux_x86_64:
	curl -Lo /tmp/protoc-3.15.6-linux-x86_64.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.15.6/protoc-3.15.6-linux-x86_64.zip
	unzip /tmp/protoc-3.15.6-linux-x86_64.zip -d /tmp/protoc-3.15.6-linux-x86_64
	chmod +x /tmp/protoc-3.15.6-linux-x86_64/bin/protoc
	mv /tmp/protoc-3.15.6-linux-x86_64/bin/protoc /usr/bin/protoc
test:
	go test -v ./tests -race
compile:
	protoc -I=protobuf --go_out=proto --go-grpc_out=proto protobuf/*.proto
