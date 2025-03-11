package main

import (
	"fmt"
	"goprototype2/internal/file"
)

func main() {
	file1 := file.NewFile("File1")
	file2 := file.NewFile("File2")
	file3 := file.NewFile("File3")
	folder1 := file.NewFolder(
		"Folder1",
		file1,
	)
	folder2 := file.NewFolder(
		"Folder2",
		folder1, file2, file3,
	)
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")
	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
