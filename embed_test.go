package main_test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string //boleh membuat lebih dari 1

//embed file ke dalam type data string
func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}

//go:embed Arabica.jpg
var Arabica []byte

//Embed file ke []byte
func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo_new.jpg", Arabica, fs.ModePerm)
	if err!= nil {
		panic(err)
	}
}

//go:embed files/one.txt
//go:embed files/two.txt
//go:embed files/three.txt
var files embed.FS

//embed multiple files
func TestMultipleFiles(t *testing.T) {
	byteOne, _ := files.ReadFile("files/one.txt")
	fmt.Println(string((byteOne)))

	byteTwo, _ := files.ReadFile("files/two.txt")
	fmt.Println(string((byteTwo)))

	byteThree, _ := files.ReadFile("files/three.txt")
	fmt.Println(string((byteThree)))
}

//go:embed files/*.txt
var path embed.FS

//path matcher
func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}