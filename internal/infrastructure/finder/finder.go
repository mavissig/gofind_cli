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

func IsSymlink(path string) (bool, string) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return false, ""
	}
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		linc, err := os.Readlink(path)
		if err != nil {
			return false, ""
		}
		return true, linc
	}
	return false, ""
}

func (f *Find) output(templatePath string) {
	prefix := templatePath + "/"
	files, err := os.ReadDir(templatePath)
	if err != nil {
		log.Println(err)
		return
	}

	tmpPath := ""
	if templatePath != "." {
		tmpPath = templatePath
	}

	for _, file := range files {

		if f.f && !file.IsDir() {

			if f.ext && path.Ext(file.Name()) == f.templateExt {
				fmt.Println(prefix + file.Name())
				continue
			}
			if !f.ext {
				check, pathLink := IsSymlink(tmpPath + file.Name())
				if check {
					fmt.Printf("%s -> %s\n", file.Name(), pathLink)
					continue
				}
				fmt.Println(prefix + file.Name())
			}
		}
		//} else if f.f && !file.IsDir() && f.ext && path.Ext(file.Name()) == f.templateExt {
		//	fmt.Println(prefix + file.Name())
		//}

		if f.d && file.IsDir() {
			fmt.Println(prefix + file.Name())
		}

		if file.IsDir() {
			f.output(prefix + file.Name())
		}

		if f.sl && !f.f {
			check, pathLink := IsSymlink(tmpPath + file.Name())
			if check {
				fmt.Printf("%s -> %s\n", file.Name(), pathLink)
				continue
			}
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
