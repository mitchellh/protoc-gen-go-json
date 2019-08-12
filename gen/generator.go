package gen

import (
	"fmt"
	"go/format"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	gen "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/generator"
)

type generator struct {
	reg *descriptor.Registry
}

// New returns a generator which generates Go files that implement
// json.Marshaler and json.Unmarshaler for the declared message types.
func New(reg *descriptor.Registry) gen.Generator {
	return &generator{reg: reg}
}

// Generator implements gen.Generator from protoc-gen-grpc-gateway
func (g *generator) Generate(targets []*descriptor.File) ([]*plugin.CodeGeneratorResponse_File, error) {
	var files []*plugin.CodeGeneratorResponse_File
	for _, file := range targets {
		glog.V(1).Infof("Processing %s", file.GetName())
		code, err := g.generate(file)
		if err != nil {
			return nil, err
		}

		formatted, err := format.Source([]byte(code))
		if err != nil {
			glog.Errorf("%v: %s", err, code)
			return nil, err
		}

		name := file.GetName()
		ext := filepath.Ext(name)
		base := strings.TrimSuffix(name, ext)
		output := fmt.Sprintf("%s.pb.json.go", base)
		files = append(files, &plugin.CodeGeneratorResponse_File{
			Name:    proto.String(output),
			Content: proto.String(string(formatted)),
		})
		glog.V(1).Infof("Will emit %s", output)
	}

	return files, nil
}

func (g *generator) generate(file *descriptor.File) (string, error) {
	return applyTemplate(file)
}
