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
	var s service.Stringer
	str := service.Read(fileName, s) //читаем строку из файла fileName
	fmt.Println(str)
	unMasked := []byte(str)
	key := []byte(keyLine)
	num := mask.Search(unMasked, key) //получаем номер символа, с которого начинается маскировка
	masked := "\n"
	masked += mask.Mask(num, unMasked, maskSym) //получаем новую строку с замаскированной ссылкой
	service.Write(fileName, masked, s)          //записываем строку в файл
}
