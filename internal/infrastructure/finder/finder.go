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

	if f.d {
		fmt.Println(templatePath)
	}

	for _, file := range files {

		if f.f && !file.IsDir() {

			if f.ext && path.Ext(file.Name()) == f.templateExt {
				fmt.Println(prefix + file.Name())
				continue
			}
			if !f.ext {
				check, pathLink := isSymlink(prefix + file.Name())
				if check {
					fmt.Printf("%s -> %s\n", prefix+file.Name(), pathLink)
					continue
				}
				fmt.Println(prefix + file.Name())
			}
		}

		if file.IsDir() {
			f.output(prefix + file.Name())
		}

		check, pathLink := isSymlink(templatePath + file.Name())
		if f.sl && check {
			fmt.Printf("%s -> %s\n", templatePath+file.Name(), pathLink)
			continue
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
