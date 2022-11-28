package app

type CPU struct {
	a       Register
	b       Register
	carry   Register
	pc      Register
	decoder Decoder
	rom     Rom
}

func init() {

}

func fetch() {

}

func decode() {

}

func execute() {

}

func NewCPU() {

}

func Run() {
	for {
		fetch()
		decode()
		execute()
	}
}
