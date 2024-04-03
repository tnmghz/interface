package service

import (
	"bufio"
	"log"
	"os"
)

type MyFile struct {
	name string
	Text string
}

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
