package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//open file for read access
	file, err := os.Open(".gitignore")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 1000)
	//read data into byte array and assign no. of bytes to count
	count, err := file.Read(data)
	defer file.Close() //closing the file after all operations are done

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	// func Chdir(dir string) error // Chdir changes the current working directory to the named directory. If there is an error, it will be of type *PathError.
	//func Chmod(name string, mode FileMode) error //Chmod changes the mode of the named file to mode. If the file is a symbolic link, it changes the mode of the link's target. If there is an error, it will be of type *PathError.
	//func Chown(name string, uid, gid int) error //Chown changes the numeric uid and gid of the named file.
}
