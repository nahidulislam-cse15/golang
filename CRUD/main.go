package main

import (
	"database/sql"
	"fmt"

	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
var temp *template.Template

type Product struct {
	ID                int
	Name              string
	Quantity_in_stock int
	Price             float32
	Description       string
}

func init() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/store")
	if err != nil {
		fmt.Println(err.Error())
	}
	//	defer db.Close()
	fmt.Println("db connection succesful")
}

func main() {
	temp, _ = template.ParseGlob("template/*.gohtml")
	http.HandleFunc("/", home)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/successful", succesful)
	http.ListenAndServe(":8080", nil)
}
func home(w http.ResponseWriter, r *http.Request) {

	sql := "select * from products"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Quantity_in_stock, &p.Price); err != nil {
			fmt.Println(err)
		}
		products = append(products, p)
		//fmt.Printf("id %d name is %s\n", id, name)

	}
	temp.ExecuteTemplate(w, "select.gohtml", products)
}
func succesful(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "insert/ update/delete succesfully")
}
func insert(w http.ResponseWriter, r *http.Request) {
	//temp.ExecuteTemplate(w, "template/insert.gohtml", nil)
	//fmt.Fprintf(w, "welcome to insert update")
	//temp.Execute(w, nil)
	if r.Method == "GET" {
		temp.ExecuteTemplate(w, "insert.gohtml", nil)
		return
	}
	r.ParseForm()
	name := r.FormValue("name")
	quantity_in_stock := r.FormValue("quantity_in_stock")
	unit_price := r.FormValue("price")
	if name == "" || quantity_in_stock == "" || unit_price == "" {
		fmt.Println("Error inserting row:", err)
		temp.ExecuteTemplate(w, "insert.gohtml", "Error inserting data, please check all fields.")
		return
	}
	//insert into db
	var ins *sql.Stmt
	// don't use _, err := db.Query()
	// func (db *DB) Prepare(query string) (*Stmt, error)
	ins, err = db.Prepare("INSERT INTO `store`.`products` (`name`,`quantity_in_stock` ,`unit_price`) VALUES (?, ?, ?);")
	if err != nil {
		panic(err)
	}
	defer ins.Close()
	// func (s *Stmt) Exec(args ...interface{}) (Result, error)
	res, err := ins.Exec(name, quantity_in_stock, unit_price)

	// check rows affectect???????
	rowsAffec, _ := res.RowsAffected()
	if err != nil || rowsAffec != 1 {
		fmt.Println("Error inserting row:", err)
		temp.ExecuteTemplate(w, "insert.gohtml", "Error inserting data, please check all fields.")
		return
	}
	lastInserted, _ := res.LastInsertId()
	rowsAffected, _ := res.RowsAffected()
	fmt.Println("ID of last row inserted:", lastInserted)
	fmt.Println("number of rows affected:", rowsAffected)
	temp.ExecuteTemplate(w, "insert.gohtml", "Product Successfully Inserted")

}
func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to home update")
}
func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to delete")
}
