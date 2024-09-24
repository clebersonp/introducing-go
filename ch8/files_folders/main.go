package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.Open("ch8/files_folders/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// get the file info
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(stat, ", size:", stat.Size())

	// read the file
	bs := make([]byte, stat.Size())
	if _, err = file.Read(bs); err != nil {
		panic(err)
	}

	str := string(bs)
	fmt.Println("File", stat.Name(), ", Content:", str)

	// shorter way to read file content as string
	bstr, err := os.ReadFile("ch8/files_folders/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Reading file content with os.ReadFile function. Content:", string(bstr))

	nfile, err := os.Create("ch8/files_folders/test_2.txt")
	if err != nil {
		panic(err)
	}
	defer nfile.Close()

	nfile.WriteString("adding text to the file")

	// get contents of a directory.
	dir, err := os.Open("ch8/files_folders")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	fileInfos, err := dir.ReadDir(-1)
	if err != nil {
		panic(err)
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	// To recursively walk a folder (read the folder's contents, all the subfolders, all the sub-subfolders)
	// There's a Walk function provided in the path/filepath package
	filepath.Walk("ch8", func(path string, info fs.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
