package service

type MyFile struct {
	name string
	Text string
}
type producer interface {
	Read()
}
type presenter interface {
	Write()
}
type Service struct {
	prod producer
	pres presenter
}

func (s Service) NewMyFile(n string) *MyFile {
	x := new(MyFile)
	x.name = n
	return x
}
