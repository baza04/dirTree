package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	isF  bool
	path string
)

func main() {
	readArgs()
	fmt.Printf("path: '%s', -f: %v\n", path, isF)
	dirTree()
}

func readArgs() {
	path = "."
	for index, arg := range os.Args[1:] {
		if index == 0 {
			path = arg
		}
		if index > 0 && arg == "-f" {
			isF = true
			return
		}
	}
}

func dirTree() {
	// child := "├───"

	ReadDir(path)

}

func ReadDir(path string) {
	content, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	PrintC(content, path)
}

func PrintC(content []os.FileInfo, curPath string) {
	for index, object := range content {
		isLast := index == len(content)-1
		if object.IsDir() {
			// fmt.Println("├───", object.Name())
			PrintFile(curPath, object, isLast)
			// curPath + "/" + object.Name()
			ReadDir(curPath + "/" + object.Name())
		} else {
			PrintFile(curPath, object, isLast)
		}
	}

}

func PrintFile(curPath string, object os.FileInfo, isLast bool) {
	pref := "├───"
	indent := countIndent(curPath, object.IsDir())
	if isLast {
		pref = "└───"
	}
	data := object.Name()
	if !object.IsDir() {
		data += "(" + strconv.Itoa(object.Size()) + "b)"
	}
	if curPath != path {
		fmt.Printf("%s%s%s\n", indent, pref, data)
	} else {
		fmt.Printf("%s%s\n", pref, data)
	}
}

func countIndent(curPath string, isDir bool) (indent string) {
	curPath = strings.Replace(curPath, path, "", 1)
	for _, char := range curPath {
		if char == '/' {
			indent += "  "
		}
	}

	if indent != "" && !isDir {
		indent = "  │" + indent[2:]
	}
	return
}
