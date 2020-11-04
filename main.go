package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

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

type parameters struct {
	isF       bool
	startPath string
	out       io.Writer
}

type stats struct {
	isDir   bool
	isLast  bool
	size    int64
	name    string
	indent  string
	curPath string
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	p := parameters{
		isF:       printFiles,
		out:       out,
		startPath: path,
	}

	return p.getInfo(p.startPath, stats{isLast: true})
}

func (p *parameters) getInfo(path string, parent stats) error {
	content, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	if !p.isF {
		content = removeFiles(content)
	}

	indent := getIndent(parent.isLast, parent.indent)

	for index, object := range content {
		stat := stats{ // fill stats
			isDir:   object.IsDir(),
			isLast:  index == len(content)-1,
			size:    object.Size(),
			curPath: path + "/" + object.Name(),
			name:    object.Name(),
			indent:  indent,
		}
		// fmt.Println("getInfo:", stat.curPath, p.startPath)

		p.PrintFile(stat)
		if stat.isDir {
			if err := p.getInfo(stat.curPath, stat); err != nil {
				return err
			}
		}
	}
	return nil
}

func removeFiles(allContent []os.FileInfo) (onlyDirs []os.FileInfo) {
	for _, object := range allContent {
		if object.IsDir() {
			onlyDirs = append(onlyDirs, object)
		}
	}
	return
}

func getIndent(isLast bool, indent string) string {
	if !isLast {
		indent += "|"
	}
	return indent + "\t"
}

func (p *parameters) PrintFile(stat stats) {
	indent := stat.indent
	// indent := stat.countIndent(p.startPath)
	pref := stat.getPreffix()
	info := stat.getNameAndSize()

	if stat.curPath != p.startPath {
		fmt.Fprintf(p.out, "%s%s%s\n", indent, pref, info)
	} else {
		fmt.Fprintf(p.out, "%s%s\n", pref, info)
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
