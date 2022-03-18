package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// getting working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dir)
	createFile("test.txt", "Golang is best programming Language")
	//file exist & info
	fileInformations("test.txt")
	makeDirectoryOutside(dir)
	//rename 
	os.Rename("test.txt","test01.txt")

}

//create file
func createFile(fileName, content string) {
	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer posf.Close()

	n, err := posf.Write([]byte(content))
	if err != nil {
		fmt.Println(err.Error())

	}
	fmt.Println(n)

}

////file exist & info
func fileInformations(fileName string) {

	fi, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("dir status=", fi.IsDir())
	fmt.Println("Name=", fi.Name())
	fmt.Println("Size= ", fi.Size())
	fmt.Printf("Modification Time=")
	fmt.Println(fi.ModTime().Date())
}

//making directory
func makeDirectory(dirName string) {
	err := os.Mkdir(dirName, 777)
	if err != nil {
		fmt.Printf("Error creating directory")
	}
}
func makeDirectoryOutside(dir string){
	
base:=filepath.Base(dir)//base returns last element of path
fmt.Println(base)
relativePath:=filepath.Join("testing")
fmt.Println(relativePath)
absolutepath,err:=filepath.Abs("testing")
if err != nil {
	fmt.Println(err.Error())
}
fmt.Println(absolutepath)
newPath:=filepath.Join(absolutepath,"..","..","..","newFolder2")
fmt.Println(newPath)
//makeDirectory(newPath)
// external url in backtic
//makeDirectory(`D:\Externaltest`)
}