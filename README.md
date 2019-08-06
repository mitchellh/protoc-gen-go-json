# protoc-gen-go-json

This is a plugin for the Google Protocol Buffers compiler
[protoc](https://github.com/protocolbuffers/protobuf) that generates
code to implement [json.Marshaler](https://golang.org/pkg/encoding/json/#Marshaler)
and [json.Unmarshaler](https://golang.org/pkg/encoding/json/#Unmarshaler)
using [jsonpb](https://godoc.org/github.com/golang/protobuf/jsonpb).

This enables Go-generated protobuf messages to be embedded directly within
other structs and encoded with the standard JSON library.
