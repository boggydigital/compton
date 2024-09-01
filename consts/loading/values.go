package loading

type Loading int

const (
	Lazy Loading = iota
	Eager
)

var loadingStrings = map[Loading]string{
	Lazy:  "lazy",
	Eager: "eager",
}

func (l Loading) String() string {
	return loadingStrings[l]
}
