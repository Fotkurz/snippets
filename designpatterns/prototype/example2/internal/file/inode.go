package file

type Inode interface {
	Print(string)
	Clone() Inode
}
