package main

type Foo struct {
	verbosity int
}

type option func(*Foo) option

// Option sets the options specified.
// It returns the previous value of the last argument.
func (f *Foo) Option(opts ...option) (previous option) {
	for _, opt := range opts {
		previous = opt(f)
	}
	return previous
}

// Verbosity sets Foo's verbosity level to v.
func Verbosity(v int) option {
	return func(f *Foo) option {
		previous := f.verbosity
		f.verbosity = v
		return Verbosity(previous)
	}
}

func DoSomethingVerbosely(foo *Foo, verbosity int) {
	// Could combine the next two lines,
	// with some loss of readability.
	prev := foo.Option(Verbosity(verbosity))
	defer foo.Option(prev)
	// ... do some stuff with foo under high verbosity.
}

func main() {

}
