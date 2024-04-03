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
