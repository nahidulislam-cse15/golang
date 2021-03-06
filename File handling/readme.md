Date:18 March 2022
# File Hahndling
The operating system provides us an interface to the device in the form of a file.
os package file handling/manipulation 
manipultion -add,delete , update etc

## Getting working directory
```
dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dir)
```
## making new directory/folder
```
unc makeDirectory(dirName string){
	err:=os.Mkdir(dirName, 777)
	if err != nil {
		fmt.Printf("Error creating directory")
	}
}

``` 
## making new directory/folder outside the working directory
```
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
```
## Creating File 
if exist then overwrite it else creates new file
```
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
```
## file info
```
func fileInformations(fileName string) {
	
	fi,err:=os.Stat(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("dir status=",fi.IsDir())
	fmt.Println("Name=",fi.Name())
	fmt.Println("Size= ",fi.Size())
    fmt.Printf("Modification Time=")
	fmt.Println( fi.ModTime().Date())
}
```
## Rename 
```
	os.Rename("test.txt","test01.txt")//(	os.Rename(oldpath ,newpath)

```
return err

Date :19 March,2022
## Delete a file
```
os.Remove(filename)
```
## File permission Mode check & change
6-rw, 7-rwe

chown not suported by windows
Date -20 March 2022
## Seek Positions in File
The function File.Seek() Seek sets the offset for the next Read or Write on file to offset, interpreted according to whence: 0 means relative to the origin of the file, 1 means relative to the current offset, and 2 means relative to the end. It returns the new offset and an error, if any.

## Write Bytes to a File
You can write using just the os package which is needed already to open the file. Since all Go executables are statically linked binaries, every package you import increases the size of your executable. Other packages like io, ioutil, and bufio provide some more help, but they are not necessary.

## Quick Write to File
The ioutil package has a useful function called WriteFile() that will handle creating/opening, writing a slice of bytes, and closing. It is useful if you just need a quick way to dump a slice of bytes to a file.
## Use Buffered Writer
The bufio package lets you create a buffered writer so you can work with a buffer in memory before writing it to disk. This is useful if you need to do a lot manipulation on the data before writing it to disk to save time from disk IO. It is also useful if you only write one byte at a time and want to store a large number in memory before dumping it to file at once, otherwise you would be performing disk IO for every byte. That puts wear and tear on your disk as well as slows down the process.

## Read up to n Bytes from File
The os.File type provides a couple basic functions. The io, ioutil, and bufio packages provided additional functions for working with files.
## Quick Read Whole File to Memory
```
 data, err := ioutil.ReadFile("test.txt")
```
## Use Buffered Reader
Creating a buffered reader will store a memory buffer with some of the contents. A buffered reader also provides some more functions that are not available on the os.File type or the io.Reader. Default buffer size is 4096 and minimum size is 16.

##  Read with a Scanner
Scanner is part of the bufio package. It is useful for stepping through files at specific delimiters. Commonly, the newline character is used as the delimiter to break up a file by lines. In a CSV file, commas would be the delimiter. The os.File can be wrapped in a bufio.Scanner just like a buffered reader. We call Scan() to read up to the next delimiter, and then use Text() or Bytes() to get the data that was read.

The delimiter is not just a simple byte or character. There is actually a special function you have to implement that will determine where the next delimiter is, how far forward to advance the pointer, and what data to return. If no custom SplitFunc is provided, it defaults to ScanLines which will split at every newline character. Other split functions included in bufio are ScanRunes, and ScanWords.
```
// To define your own split function, match this fingerprint
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

// Returning (0, nil, nil) will tell the scanner
// to scan again, but with a bigger buffer because
// it wasn't enough data to reach the delimiter
```

## Archive(Zip) Files

## Extract(Unzip) Archived Files
## Compress a File
## Uncompress a File
## Temporary Files and Directories
The ioutil package provides two functions: TempDir() and TempFile(). It is the callers responsibility to delete the temporary items when done. The only benefit these functions provide is that you can pass it an empty string for the directory, and it will automatically create the item in the system's default temporary folder (/tmp on Linux). Since os.TempDir() function that will return the defauly system temporary directory.

## Downloading a File Over HTTP
## Hashing and Checksums
