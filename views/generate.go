// +build ignore

package main

import (
	"github.com/shurcooL/httpfs/filter"
	"github.com/shurcooL/httpfs/union"
	"github.com/shurcooL/vfsgen"
	"net/http"
)

func main() {
	fs := union.New(map[string]http.FileSystem{
		"/static":    filter.Skip(http.Dir("static"), filter.FilesWithExtensions(".go")),
		"/templates": http.Dir("templates"),
	})
	if err := vfsgen.Generate(fs, vfsgen.Options{
		PackageName:  "static",
		Filename:     "static/assets.go",
		VariableName: "Assets",
	}); err != nil {
		panic(err.Error())
	}
}
