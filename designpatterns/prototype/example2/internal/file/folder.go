package file

import "fmt"

type Folder struct {
	name     string
	children []Inode
}

func NewFolder(name string, children ...Inode) Inode {
	f := &Folder{
		name:     name,
		children: children,
	}
	return f
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.name)

	for _, inode := range f.children {
		inode.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() Inode {
	cloneFolder := &Folder{name: f.name, children: f.children}
	var tempChildren []Inode
	for _, inode := range f.children {
		copy := inode.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
