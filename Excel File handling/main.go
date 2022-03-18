package main

import (
	"fmt"
   _ "image/gif"
    _ "image/jpeg"
    _ "image/png"

	"github.com/xuri/excelize"
)

func main() {
	//create excel file
	f := excelize.NewFile()
	// Create a new worksheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set the active worksheet of the workbook.
	f.SetActiveSheet(index)
	// Save the spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
	readingExcel()
	addChart()
	addImage()

}

//reading excel file
func readingExcel() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from the cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

//add chart
func addChart() {

	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	index := f.NewSheet("Sheet2")
	f.SetActiveSheet(index)

	for k, v := range categories {
		f.SetCellValue("Sheet2", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet2", k, v)
	}
	if err := f.AddChart("Sheet2", "E1", `{
        "type": "col3DClustered",
        "series": [
        {
            "name": "Sheet2!$A$2",
            "categories": "Sheet2!$B$1:$D$1",
            "values": "Sheet2!$B$2:$D$2"
        },
        {
            "name": "Sheet2!$A$3",
            "categories": "Sheet2!$B$1:$D$1",
            "values": "Sheet2!$B$3:$D$3"
        },
        {
            "name": "Sheet2!$A$4",
            "categories": "Sheet2!$B$1:$D$1",
            "values": "Sheet2!$B$4:$D$4"
        }],
        "title":
        {
            "name": "Fruit 3D Clustered Column Chart"
        }
    }`); err != nil {
		fmt.Println(err)
		return
	}
	// Save the spreadsheet by the given path.
	if err := f.SaveAs("Book2.xlsx"); err != nil {
		fmt.Println(err)
	}
}
func addImage() {
    f := excelize.NewFile()
    if err := f.SaveAs("Book3.xlsx"); err != nil {
		fmt.Println(err)
	}
	f, err := excelize.OpenFile("Book3.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Insert a picture.
	if err := f.AddPicture("Sheet1", "A2", "image.png", ""); err != nil {
		fmt.Println(err)
	}
	// Insert a picture to worksheet with scaling.
	if err := f.AddPicture("Sheet1", "D2", "image.jpg", `{
        "x_scale": 0.5,
        "y_scale": 0.5
    }`); err != nil {
		fmt.Println(err)
	}
	// Insert a picture offset in the cell with printing support.
	if err := f.AddPicture("Sheet1", "H2", "image.gif", `{
        "x_offset": 15,
        "y_offset": 10,
        "print_obj": true,
        "lock_aspect_ratio": false,
        "locked": false
    }`); err != nil {
		fmt.Println(err)
	}
	// Save the spreadsheet with the origin path.
	if err = f.Save(); err != nil {
		fmt.Println(err)
	}
	if err = f.Close(); err != nil {
		fmt.Println(err)
	}
}
