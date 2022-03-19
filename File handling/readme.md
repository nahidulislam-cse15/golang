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



