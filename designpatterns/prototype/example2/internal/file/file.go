package file

import "fmt"

type File struct {
	name string
}

func NewFile(name string) Inode {
	return &File{name: name}
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) Clone() Inode {
	return &File{name: f.name + "_clone"}
}
