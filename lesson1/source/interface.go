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

type Stream struct {
	Address string
	Data    string
}

func (f Stream) Read() string {
	return f.Data
}

type Reader interface {
	Read() string
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

func Process(reader Reader) {
	fmt.Println(reader.Read())
}
