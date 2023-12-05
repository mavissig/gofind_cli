package finder

import (
	"fmt"
	"log"
	"os"
	"path"
)

type Find struct {
	flags
}

type flags struct {
	sl, d, f, ext bool
	templateExt   string
	templatePath  string
}

func (f *Find) parsingFlags(flags map[string]bool) {
	f.sl = flags["sl"]
	f.d = flags["d"]
	f.f = flags["f"]
	f.ext = flags["ext"]
}

func (f *Find) parsingArgs(args []string) {
	for i, arg := range args {
		if f.ext && i == 0 {
			f.templateExt = "." + arg
			continue
		}
		f.templatePath = arg
	}
}

func (f *Find) output(templatePath string) {
	prefix := templatePath + "/"
	files, err := os.ReadDir(templatePath)
	if err != nil {
		log.Println(err)
		return
	}
	for _, file := range files {
		if f.f && !file.IsDir() && !f.ext {
			fmt.Println(prefix + file.Name())
		}
		if f.f && !file.IsDir() && f.ext && path.Ext(file.Name()) == f.templateExt {
			fmt.Println(prefix + file.Name())
		}
		if f.d && file.IsDir() {
			fmt.Println(prefix + file.Name())
		}
		if file.IsDir() {
			f.output(prefix + file.Name())
		}
	}
}

func (f *Find) Do(flags map[string]bool, args []string) {
	f.parsingFlags(flags)
	f.parsingArgs(args)
	f.output(f.templatePath)
}

func New() *Find {
	return &Find{}
}
