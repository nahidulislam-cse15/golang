package main

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// getting working directory
	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(dir)
	// createFile("test.txt", "Golang is best programming Language")
	// //file exist & info
	// fileInformations("test.txt")
	// makeDirectoryOutside(dir)
	// //rename
	// os.Rename("test.txt", "test01.txt")

	// truncateFile("test01.txt")
	// //	fileInformations("test.txt")
	// fileInformations("test01.txt")
	// createFile("test.txt", "Golang is best programming Language")

	//DeleteFile("test.txt")
	// CheckPermissions("test.txt")
	// change("test.txt")
	//copyFile("test_copy.txt", "test.txt")
	//copy file
	// originalFile, err := os.Open("test.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// defer originalFile.Close()
	// newFile, err := os.Create("test_copy.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// defer newFile.Close()
	// bytesWritten, err := io.Copy(newFile, originalFile)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Printf("Copied %d bytes.", bytesWritten)

	//// //fmt.Println(byteswritten)
	seekFile()
	writeFile()
	//Quick Write to File
	quickWrite()
	bufferWrite()
	readFile()
	bufferReader()
		scannerReader()
	archiveFile()
	extractZip()
		compressZip()
		uncompressZip()
	temporaryFile()
		//downloadOverHttp()
	hashing()

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
	bytesRead2, err := bufferRd.ReadBytes('\b')
	if err != nil {
		fmt.Println("error in read")
	}
	fmt.Printf("Bytes read:%s", bytesRead2)
	//fmt.Println("Data read:",byteSlice)
	bytesreadstring, err := bufferRd.ReadString('\n')
	if err != nil {
		fmt.Println("error in read")
	}
	fmt.Printf("Bytes read:%s", bytesreadstring)
}
func scannerReader() {
	file, err := os.Open("test01.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	// fmt.Println(token)
	//success := scanner.Scan()
	//fmt.Println(success)
	for scanner.Scan() {
		// Get data from scan with Bytes() or Text()
		fmt.Println(scanner.Text())
		fmt.Println(scanner.Bytes())
	}
	//fmt.Println("First word found:", scanner.Text())
}
func archiveFile() {

	zfile, err := os.Create("test.zip")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer zfile.Close()
	zipWriter := zip.NewWriter(zfile)
	filesToArchive := []string{"test.txt", "test01.txt"}

	// Create and write files to the archive, which in turn
	// are getting written to the underlying writer to the
	// .zip file we created at the beginning
	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file)
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			// Get data from scan with Bytes() or Text()
			fmt.Println(scanner.Text())
			_, err = fileWriter.Write([]byte(scanner.Text()))
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	err = zipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}

}
func extractZip() {
	//create reader
	zipReader, err := zip.OpenReader("test.zip")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer zipReader.Close()
	// Iterate through each file/dir found in
	for _, file := range zipReader.Reader.File {
		// Open the file inside the zip archive
		// like a normal file
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()
		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)
		// Extract the item (or create directory)
		if file.FileInfo().IsDir() {
			// Create directories to recreate directory
			// structure inside the zip archive. Also
			// preserves permissions
			log.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			// Extract regular file since not a directory
			log.Println("Extracting file:", file.Name)
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}

func compressZip() {
	outputFile, err := os.Create("testc.txt.gz")
	if err != nil {
		fmt.Println(err.Error())
	}
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()
	_, err = gzipWriter.Write([]byte("hello world"))
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Compressed data written to file.")
}
func uncompressZip() {
	gzipFile, err := os.Open("testc.txt.gz")
	if err != nil {
		fmt.Println(err.Error())
	}
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer gzipReader.Close()
	outputWriter, err := os.Create("uncompress.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer outputWriter.Close()

	_, err = io.Copy(outputWriter, gzipReader)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func temporaryFile() {
	// Create a temp dir in the system default temp folder
	tempDirPath, err := ioutil.TempDir("", "tempDir")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Temp dir created:", tempDirPath)
	// Create a file in new temp directory
	tempFile, err := ioutil.TempFile(tempDirPath, "TempFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Temp file created:", tempFile.Name())
	// Close file
	err = tempFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	// Delete the resources we created
	err = os.Remove(tempFile.Name())
	if err != nil {
		fmt.Println(err.Error())
	}
	err = os.Remove(tempDirPath)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func downloadOverHttp() {
	newFile, err := os.Create("devdungeon.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// HTTP GET request devdungeon.com
	url := "https://www.devdungeon.com/archive"
	response, err := http.Get(url)
	defer response.Body.Close()

	// Write bytes from HTTP response to file.
	// response.Body satisfies the reader interface.
	// newFile satisfies the writer interface.
	// That allows us to use io.Copy which accepts
	// any type that implements reader and writer interface
	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)

}
func hashing() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Hash the file and output results
	fmt.Printf("Md5: %x\n\n", md5.Sum(data))
	fmt.Println(md5.Sum(data))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
	fmt.Println(sha1.Sum(data))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
	fmt.Println(sha256.Sum256(data))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))
	fmt.Println(sha512.Sum512(data))
	//custom hash
	// Open file for reading
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    // Create new hasher, which is a writer interface
    hasher := sha1.New()
    _, err = io.Copy(hasher, file)
    if err != nil {
        log.Fatal(err)
    }

    // Hash and print. Pass nil since
    // the data is not coming in as a slice argument
    // but is coming through the writer interface
    sum := hasher.Sum(nil)
    fmt.Printf("Md5 checksum: %x\n", sum)
}
