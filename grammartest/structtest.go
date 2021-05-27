package grammartest

func NewFactory() *Factory {
	return &Factory{
		Name: "FF",
		Sex:  1,
	}
}

type Factory struct {
	Name string
	Sex  int
}
