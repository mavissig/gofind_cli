package finder

import "fmt"

type Find struct {
	flags
}

type flags struct {
	sl, d, f bool
}

func (f *Find) parsingFlags() {

}

func (f *Find) Do() {
	fmt.Println("used find util")
}

func New() *Find {
	return &Find{}
}
