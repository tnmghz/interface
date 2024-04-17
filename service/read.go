package service

import (
	"bufio"
	"log"
	"os"
)

func (mf *MyFile) Read() {
	file, err := os.Open(mf.name)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mf.Text = scanner.Text()
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
func Read(name string, s Stringer) string {
	var ms myString
	str := ms.stringRead(name)
	return str
}

func (ms myString) stringRead(name string) string {
	var serv Service
	x := serv.NewMyFile(name)
	x.Read()
	s := x.Text
	return s
}

func (m *mockString) stringRead(name string) string {
	args := m.Called(name)
	return args.String(0)
}
