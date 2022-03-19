package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
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
	os.Rename("test.txt", "test01.txt")

	truncateFile("test01.txt")
	//	fileInformations("test.txt")
	fileInformations("test01.txt")
	createFile("test.txt", "Golang is best programming Language")

	//DeleteFile("test.txt")
	CheckPermissions("test.txt")
	change("test.txt")
	//copyFile("test_copy.txt", "test.txt")
	originalFile, err := os.Open("test.txt")
    if err != nil {
        fmt.Println(err.Error())
    }
    defer originalFile.Close()
	newFile, err := os.Create("test_copy.txt")
    if err != nil {
		fmt.Println(err.Error())
    }
    defer newFile.Close()
	bytesWritten, err = io.Copy(newFile, originalFile)
    if err != nil {
		fmt.Println(err.Error())
    }
    fmt.Printf("Copied %d bytes.", bytesWritten)
	
	//fmt.Println(byteswritten)

}

//create file
func createFile(fileName, content string) {
	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(*posf)
	//log.Println(*posf)
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
	fmt.Println("Permissions:", fi.Mode())
	fmt.Printf("System interface type: %T\n", fi.Sys())
	fmt.Printf("System info: %+v\n\n", fi.Sys())
}

//making directory
func makeDirectory(dirName string) {
	err := os.Mkdir(dirName, 777)
	if err != nil {
		fmt.Printf("Error creating directory")
	}
}
func makeDirectoryOutside(dir string) {

	base := filepath.Base(dir) //base returns last element of path
	fmt.Println(base)
	relativePath := filepath.Join("testing")
	fmt.Println(relativePath)
	absolutepath, err := filepath.Abs("testing")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(absolutepath)
	newPath := filepath.Join(absolutepath, "..", "..", "..", "newFolder2")
	fmt.Println(newPath)
	//makeDirectory(newPath)
	// external url in backtic
	//makeDirectory(`D:\Externaltest`)
}

//truncating file
func truncateFile(fileName string) {
	err := os.Truncate(fileName, 20)
	if err != nil {
		fmt.Println(err.Error())
	}

}

//delete file
func DeleteFile(fileName string) {
	if err := os.Remove(fileName); err != nil {
		fmt.Println(err.Error())
	}
}

//check permissions
func CheckPermissions(fileName string) {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Println("Error: Write permission denied.")
		}
	}
	file.Close()

	// Test read permissions
	file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Println("Error: Read permission denied.")
		}
	}
	file.Close()
}

//change permissions
func change(fileName string) {
	err := os.Chmod(fileName, 0777)
	if err != nil {
		fmt.Println(err)
	}

	// Change ownership
	err = os.Chown(fileName, os.Getuid(), os.Getgid())
	if err != nil {
		fmt.Println(err)
	}

	// Change timestamps
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes(fileName, lastAccessTime, lastModifyTime)
	if err != nil {
		fmt.Println(err)
	}
}

//copying file
func copyFile(oldfile string) {

	
}
