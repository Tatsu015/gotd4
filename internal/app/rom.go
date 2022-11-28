package app

type ROM struct{}

func NewROM() ROM {
	return ROM{}
}

func (r *ROM) Fetch() string {
	return "test" // TODO
}
