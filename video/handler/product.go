package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"video/db"
)

type Product struct {
	Id           int
	Name         string
	Brand        string
	Price        float32
	Description  string
	Created_date time.Time
	Status       bool
}

func ShowproductHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("show product list handler runing ...")
	tmpl.ExecuteTemplate(w, "Productlist.html", nil)

}

func AddproductHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add product handler runing ...")
	db := db.DbConn()

	email := getUserName(r)
	fmt.Println("login user..", email)
	p1 := person{
		email,
	}

	if r.Method == "POST" {
		name := r.FormValue("name")
		brand := r.FormValue("brand")
		price := r.FormValue("price")
		description := r.FormValue("description")
		status := 1

		var id int
		currentTime := time.Now()
		err := db.QueryRow("select id from User where email = ?", email).Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("selDB = ", id)

		insForm, err := db.Prepare("INSERT INTO Products(retailer_id,name,brand,price,description,created_date,status) VALUES(?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(id, name, brand, price, description, currentTime, status)
		fmt.Println(id, name, brand, price, description, status)

		tmpl.ExecuteTemplate(w, "Productlist.html", p1)

	} else {
		tmpl.ExecuteTemplate(w, "addproduct.html", p1)
	}

}

func ProductlistHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Product list Handler runing ...")
	db := db.DbConn()

	email := getUserName(r)

	var id int
	err := db.QueryRow("select id from User where email = ?", email).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	selDB, err := db.Query("SELECT * FROM Products WHERE retailer_id=? ORDER BY id DESC", id)
	if err != nil {
		panic(err.Error())
	}
	pro := Product{}
	res := []Product{}
	for selDB.Next() {
		var id int
		var name, brand, description string
		var created_date time.Time
		var price float32
		var status bool
		err = selDB.Scan(&id, &name, &brand, &price, &description, &created_date, &status)
		if err != nil {
			panic(err.Error())
		}
		pro.Id = id
		pro.Name = name
		pro.Brand = brand
		pro.Price = price
		pro.Description = description
		pro.Created_date = created_date
		pro.Status = status
		res = append(res, pro)
	}
	fmt.Println("pro = ", pro)
	fmt.Println("res = ", res)
	tmpl.ExecuteTemplate(w, "Productlist", res)
	defer db.Close()

}
