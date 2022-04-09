package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/store")
	if err != nil {
		fmt.Println(err.Error())
	}
	//	defer db.Close()
	fmt.Println("db connection succesful")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":8080", nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to home crud\n")
	sql := "select * from products"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id, quantity_in_stock int64
			price                 float64
			name                  string
		)
		if err := rows.Scan(&id, &name, &quantity_in_stock, &price); err != nil {
			fmt.Println(err)
		}
		//fmt.Printf("id %d name is %s\n", id, name)
		fmt.Fprintf(w, "product_id:%d name:%s quantity_in_stock:%d  price:%f,\n", id, name, quantity_in_stock, price)
	}
}
func insert(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to insert")
}
func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to home update")
}
func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to delete")
}
