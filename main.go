package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type stats struct {
	isDir   bool
	isLast  bool
	size    int64
	name    string
	indent  string
	curPath string
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {

	content, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for index, object := range content {
		stat := stats{ // fill stats
			isDir:   object.IsDir(),
			isLast:  index == len(content)-1,
			curPath: path + "/" + object.Name(),
			size:    object.Size(),
			name:    object.Name(),
		}

		PrintFile(out, path, printFiles, stat)
		if stat.isDir {
			if err := dirTree(out, stat.curPath, printFiles); err != nil {
				panic(err.Error())
			}
		}
	}
	return nil
}

func PrintFile(out io.Writer, path string, printFiles bool, stat stats) {
	pref := stat.getPreffix()
	indent := ""
	// indent := countIndent(stat.curPath, stat.isDir, stat.isLast)
	info := stat.getNameAndSize()

	if stat.curPath != path {
		fmt.Fprintf(out, "%s%s%s\n", indent, pref, info)
	} else {
		fmt.Fprintf(out, "%s%s\n", pref, info)
	}
}

func (s stats) getPreffix() string {
	if s.isLast {
		return "└───"
	}
	return "├───"
}

func (s stats) getNameAndSize() string {
	if s.isDir {
		return s.name
	}
	if s.size != 0 {
		return fmt.Sprintf("%s (%db)", s.name, s.size)
	}
	return fmt.Sprintf("%s (empty)", s.name)
}

// func countIndent(curPath string, isDir, isLast bool) (indent string) {
// 	// curPath = strings.Replace(curPath, path, "", 1)
// 	for _, char := range curPath {
// 		if char == '/' {
// 			indent += "  "
// 			if /* !isDir && */ !isLast {
// 			}
// 		}
// 	}
// 	return
// }
