package main

import (
	"fmt"
)

type File struct {
	Name    string
	Content string
}

func (f File) Read() string {
	return f.Content
}

func (f File) Write() {

}

type Stream struct {
	Address string
	Data    string
}

func (f Stream) Read() string {
	return f.Data
}

func (f Stream) Write() {

}

type Reader interface {
	Read() string
}

type IOEntity interface {
	Reader
	Write()
}

func main() {
	file := File{
		Name:    "f1.txt",
		Content: "content1",
	}
	stream := Stream{
		Address: "d1.com",
		Data:    "stream_data",
	}
	Process(file)
	Process(stream)
}

func Process(entity IOEntity) {
	fmt.Println(entity.Read())
	entity.Write()

}
