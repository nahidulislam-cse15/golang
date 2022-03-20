package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
	//copy file
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
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Copied %d bytes.", bytesWritten)

	//fmt.Println(byteswritten)
	seekFile()
	writeFile()
	//Quick Write to File
	quickWrite()
	bufferWrite()
	readFile()
	bufferReader()

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
func seekFile() {
	file, _ := os.Open("test.txt")
	defer file.Close()

	// Offset is how many bytes to move
	// Offset can be positive or negative
	var offset int64 = 5

	// Whence is the point of reference for offset
	// 0 = Beginning of file
	// 1 = Current position
	// 2 = End of file
	var whence int = 0
	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved to 5:", newPosition)

	// Go back 2 bytes from current position
	newPosition, err = file.Seek(-2, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved back two:", newPosition)

	// Find the current position by getting the
	// return value from Seek after moving 0 bytes
	currentPosition, err := file.Seek(0, 1)
	fmt.Println("Current position:", currentPosition)

	// Go to beginning of file
	newPosition, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Position after seeking 0,0:", newPosition)
}

//write bytes
func writeFile() {
	// Open a new file for writing only
	file, err := os.OpenFile(
		"test.txt",
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write bytes to file
	byteSlice := []byte("is role the master!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}

//Quick Write to File
func quickWrite() {
	//Quick Write to File
	if err := ioutil.WriteFile("test2.txt", []byte("Hi\n"), 0666); err != nil {
		fmt.Println(err.Error())
	}

}
func bufferWrite() {
	f, err := os.OpenFile("test.txt", os.O_WRONLY, 777)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	//buffer writer
	bufferWr := bufio.NewWriter(f)
	//write bytesWritten
	bytesWritten, err := bufferWr.WriteString("nahid master")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(bytesWritten)
	// Check how much is stored in buffer waiting
	unflashbufffer := bufferWr.Buffered()
	fmt.Println("Bytes buffered:", unflashbufffer)
	bytesAvailable := bufferWr.Available()
	fmt.Println("Bytes available:", bytesAvailable)
	bufferWr.Flush()
	unflashbufffer = bufferWr.Buffered()
	fmt.Println("Bytes buffered:", unflashbufffer)
	bytesAvailable = bufferWr.Available()
	fmt.Println("Bytes available:", bytesAvailable)
	bufferWr.Reset(bufferWr)

	// See how much buffer is available
	bytesAvailable = bufferWr.Available()
	fmt.Println("Bytes available:", bytesAvailable)
	bufferWr = bufio.NewWriterSize(bufferWr, 8000)
	bytesAvailable = bufferWr.Available()
	fmt.Println("Bytes available:", bytesAvailable)

}

//Read up to n Bytes from File
func readFile() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	byteslice := make([]byte, 100)
	bytesread, err := file.Read(byteslice)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Bytes read:", bytesread)
	fmt.Println("Data read:", byteslice)
	//full read
	//byteslice=make([]byte,100)
	bytesread, err = io.ReadFull(file, byteslice)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Bytes read:", bytesread)
	fmt.Println("Data read:", byteslice)
	//read At least
	minBytes := 80
	bytesread2, err := io.ReadAtLeast(file, byteslice, minBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Bytes read:", bytesread2)
	fmt.Println("Data read:", byteslice)
	//read all

	file, err = os.Open("test01.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("hex %x\n", data)
	fmt.Printf("string :%s\n", data)
	fmt.Println("len data:", len(data))
	// Quick Read Whole File to Memory
	// data, err := ioutil.ReadFile("test01.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("hex %x\n", data)
	// fmt.Printf("string :%s\n", data)
	// fmt.Println("len data:", len(data))

}
func bufferReader() {
	file, err := os.Open("test01.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	bufferRd := bufio.NewReader(file)
	 // Get bytes without advancing pointer
	//  byteSlice := make([]byte, 20)
	//  byteSlice, err = bufferRd.Peek(5)
	//  if err != nil {
	// 	fmt.Println(err.Error())
	//  }
	//  fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)
	//  bytesRead,err:=bufferRd.Read(byteSlice)
	//  if err!= nil {
	// 	 fmt.Println(err.Error())
	//  }
	//  fmt.Println("Bytes read:",bytesRead)
	//  fmt.Println("Data read:",byteSlice)
	//  bytesRead1 ,err :=bufferRd.ReadByte()
	//  if err!= nil {
	// 	 fmt.Println(err.Error())
	//  }
	//  fmt.Printf("Bytes read:%c",bytesRead1)
	 bytesRead2 ,err :=bufferRd.ReadBytes('\b')
	 if err!= nil {
		 fmt.Println("error in read")
	 }
	 fmt.Printf("Bytes read:%s",bytesRead2)
	 //fmt.Println("Data read:",byteSlice)
	 bytesreadstring,err :=bufferRd.ReadString('\n')
	 if err!= nil {
		fmt.Println("error in read")
	}
	fmt.Printf("Bytes read:%s",bytesreadstring)
}
