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
