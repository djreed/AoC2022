package main

import (
	"bufio"
	"os"
)

var (
	// top-level root folder, with no parent
	root = &Directory{
		name:        "/",
		parent:      nil,
		contents:    make([]File, 0),
		directories: make([]Directory, 0),
	}

	// Pointer to the current directory -- stateful
	current *Directory
)

type Directory struct {
	name string

	parent *Directory

	contents    []File
	directories []Directory
}

type File struct {
	name string
	size int
}

func main() {
	filename := "example"
	readFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Consume input to construct our filetree
	for fileScanner.Scan() {
		// cd -> Move the directory that we are populating

		// ls -> We're getting file data, need to use to populate
	}

	// Once constructed, need to do operations on the file sizes within the tree

}

func processCommand() {

}

func moveUp() {
	current = current.parent
}

func moveDown(name string) {
	current = current.findOrCreateDirectory(name)
}

// Returns a pointer to a matching directory within this folder's children
func (parent *Directory) findOrCreateDirectory(name string) *Directory {
	for _, d := range parent.directories {
		if d.name == name {
			return &d
		}
	}

	// If not found in existing list, create new child directory
	newDirectory := Directory{name: name, parent: parent}
	parent.directories = append(parent.directories, newDirectory)

	return &newDirectory
}
