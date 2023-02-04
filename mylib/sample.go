package mylib

import "time"

func SampleFunc() (hoge string) {
	hoge = "hello"
	return
}

func Goroutin(ch chan string) {
	for {
		ch <- "2"
		time.Sleep(2 * time.Second)
	}
}