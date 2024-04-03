package main

import (
	"fmt"
	"interface/mask"
	"interface/service"
)

const (
	fileName = "data.txt"
	keyLine  = "http://"
	maskSym  = '*'
)

func main() {
	var serv service.Service
	x := serv.NewMyFile(fileName)
	x.Read() //читаем строку из файла fileName
	fmt.Println(x.Text)
	unMasked := []byte(x.Text)
	key := []byte(keyLine)
	num := mask.Search(unMasked, key) //получаем номер символа, с которого начинается маскировка
	masked := "\n"
	masked += mask.Mask(num, unMasked, maskSym) //получаем новую строку с замаскированной ссылкой
	x.Write(masked)                             //записываем строку в файл
}
