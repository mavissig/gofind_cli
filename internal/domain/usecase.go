package domain

type Finder interface {
	Do(map[string]bool, []string)
}

type UseCase struct {
	finder Finder
}

func (uc *UseCase) Find(flags map[string]bool, args []string) {
	uc.finder.Do(flags, args)
}

func New(f Finder) *UseCase {
	return &UseCase{
		finder: f,
	}
}
