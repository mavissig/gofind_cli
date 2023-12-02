package finder

type Find struct {
	flags
}

type flags struct {
	sl, d, f bool
}

func (f *Find) parsingFlags() {

}

func (f *Find) Do() {

}

func New() *Find {
	return &Find{}
}
