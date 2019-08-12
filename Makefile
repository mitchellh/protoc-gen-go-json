.PHONY: proto
proto:
	cd e2e && protoc --go_out=. --go-json_out=. e2e.proto
