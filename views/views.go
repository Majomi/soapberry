package views

import "html/template"

import "github.com/shurcooL/httpfs/html/vfstemplate"

import "github.com/shurcooL/httpfs/path/vfspath"

import "github.com/majomi/soapberry/static"

var (
	LayoutDir   string = "/templates/layouts/"
	TemplateExt string = ".gohtml"
	TemplateDir string = "/templates/"
)

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)

	files = append(files, layoutFiles()...)
	t, err := vfstemplate.ParseFiles(static.Assets, template.New(""), files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

func layoutFiles() []string {
	files, err := vfspath.Glob(static.Assets, LayoutDir+"*"+TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}

type View struct {
	Template *template.Template
	Layout   string
}
