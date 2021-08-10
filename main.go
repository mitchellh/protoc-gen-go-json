package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/mitchellh/protoc-gen-go-json/gen"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	enumsAsInts  = flag.Bool("enums_as_ints", false, "render enums as integers as opposed to strings")
	emitDefaults = flag.Bool("emit_defaults", false, "render fields with zero values")
	origName     = flag.Bool("orig_name", false, "use original (.proto) name for fields")
	allowUnknown = flag.Bool("allow_unknown", false, "allow messages to contain unknown fields when unmarshaling")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gp *protogen.Plugin) error {

		opts := gen.Options{
			EnumsAsInts:        *enumsAsInts,
			EmitDefaults:       *emitDefaults,
			OrigName:           *origName,
			AllowUnknownFields: *allowUnknown,
		}

		for _, name := range gp.Request.FileToGenerate {
			f := gp.FilesByPath[name]

			if len(f.Messages) == 0 {
				glog.V(1).Infof("Skipping %s, no messages", name)
				continue
			}

			glog.V(1).Infof("Processing %s", name)
			glog.V(2).Infof("Generating %s\n", fmt.Sprintf("%s.pb.json.go", f.GeneratedFilenamePrefix))

			gf := gp.NewGeneratedFile(fmt.Sprintf("%s.pb.json.go", f.GeneratedFilenamePrefix), f.GoImportPath)

			err := gen.ApplyTemplate(gf, f, opts)
			if err != nil {
				gf.Skip()
				gp.Error(err)
				continue
			}
		}

		return nil
	})
}
