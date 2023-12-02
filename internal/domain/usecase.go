package domain

type Finder interface {
	Do()
}

type UseCase struct {
	finder Finder
}

func (uc *UseCase) Find() {
	uc.finder.Do()
}

func New(f Finder) *UseCase {
	return &UseCase{
		finder: f,
	}
}
