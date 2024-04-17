package service

import "github.com/stretchr/testify/mock"

type MyFile struct {
	name string
	Text string
}
type producer interface {
	Read()
}
type presenter interface {
	Write(string)
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

type myString struct{}
type Stringer interface {
	stringRead(string) string
	stringWrite(string, string) string
}
type mockString struct {
	mock.Mock
}
