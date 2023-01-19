package type_func

type Factory interface {
	A(i int) int
	B(f float32) float32
}

type ImplA func(i int) int

func (ia ImplA) A(i int) int {
	return ia(i)
}

type ImplB func(f float32) float32

func (ib ImplB) B(f float32) float32 {
	return ib(f)
}

type Imp struct {
	ImplA
	ImplB
}

type Option func(i *Imp)

func WithImplA(ia ImplA) Option {
	return func(i *Imp) {
		i.ImplA = ia
	}
}

func WithImplB(ib ImplB) Option {
	return func(i *Imp) {
		i.ImplB = ib
	}
}

func NewImpl(options ...Option) Factory {
	i := &Imp{}
	for _, o := range options {
		o(i)
	}
	return i
}
