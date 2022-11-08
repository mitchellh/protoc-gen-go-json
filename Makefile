.PHONY: proto

tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

proto: tools
	mkdir -p build
	go build -o build/protoc-gen-go-json .
	export PATH=$(CURDIR)/build/:$$PATH && \
	    cd e2e && protoc --go_out=. --go-json_out=logtostderr=true,v=10,multiline=true,partial=true:. *.proto
