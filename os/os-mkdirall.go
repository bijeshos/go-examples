package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	path := "./tmp1/ab/cd"

	err := os.MkdirAll(path, os.ModeTemporary)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("directory created")
	}
}

//func NewFile(fd uintptr, name string) *File
//func Open(name string) (*File, error)
//func OpenFile(name string, flag int, perm FileMode) (*File, error)
//func (*File) Write  //Write writes len(b) bytes to the File.
//func (*File) WriteAt //WriteAt writes len(b) bytes to the File starting at byte offset off
//func (*File) WriteString //WriteString is like Write, but writes the contents of string s rather than a slice of bytes.
//func (FileMode) IsDir ¶
//func (FileMode) IsRegular ¶
