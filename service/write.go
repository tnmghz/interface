package service

import (
	"log"
	"os"
)

func (mf *MyFile) Write(s string) {
	file, err := os.OpenFile(mf.name, os.O_APPEND, 123)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(s)
	if err != nil {
		log.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
func Write(name, text string, s Stringer) string {
	var ms myString
	str := ms.stringWrite(name, text)
	return str
}

func (ms myString) stringWrite(name, text string) string {
	var serv Service
	x := serv.NewMyFile(name)
	x.Write(text)
	x.Read()
	s := x.Text
	return s
}
func (m *mockString) stringWrite(name, text string) string {
	args := m.Called(name, text)
	return args.String(0)
}
