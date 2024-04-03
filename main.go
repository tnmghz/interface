package main

import (
	"bufio"
	"fmt"
	"interface/mask"
	"log"
	"os"
)

const (
	fileName = "data.txt"
	keyLine  = "http://"
	maskSym  = '*'
)

type myFile struct {
	name string
	text string
}

func (mf *myFile) Read() {
	file, err := os.Open(mf.name)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mf.text = scanner.Text()
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
func (mf *myFile) Write(s string) {
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

func (s Service) newMyFile(n string) *myFile {
	x := new(myFile)
	x.name = n
	return x
}

func main() {
	var serv Service
	x := serv.newMyFile(fileName)
	x.Read() //читаем строку из файла fileName
	fmt.Println(x.text)
	unMasked := []byte(x.text)
	key := []byte(keyLine)
	num := mask.Search(unMasked, key) //получаем номер символа, с которого начинается маскировка
	masked := "\n"
	masked += mask.Mask(num, unMasked, maskSym) //получаем новую строку с замаскированной ссылкой
	x.Write(masked)                             //записываем строку в файл
}
