.PHONY: proto
proto:
	mkdir -p build
	go build -o build/protoc-gen-go-json .
	export PATH=$(CURDIR)/build/:$$PATH && \
	    protoc --go_out=. -I./e2e --go-json_out=logtostderr=true,v=10:. e2e/*.proto
