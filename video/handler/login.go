package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
	dbc "video/db"

	"github.com/gorilla/securecookie"
)

type person struct {
	First string
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

var tmpl = template.Must(template.ParseGlob("template/*"))

// create session with cookie functionality...
func setSession(userName string, response http.ResponseWriter) {
	fmt.Println("session created...")
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

//end session with cookie functionality

// db comman function check value into db are exists or not...
func rowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := dbc.DbConn().QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("error checking if row exists", err)
	}
	return exists
}

// end db comman function check value into db are exists or not

//login functionality ...

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login handler runing...")
	email := r.FormValue("email")
	pass := r.FormValue("psw")
	redirectTarget := "/index"
	if email != "" && pass != "" {
		fmt.Println("email = ", email)
		// .. check credentials ..
		dataexists := rowExists("SELECT * FROM User WHERE email=?", email)

		fmt.Println("dataexists = ", dataexists)
		if dataexists {
			setSession(email, w)
			redirectTarget = "/dashboard"
		}

	}
	http.Redirect(w, r, redirectTarget, 302)
}

//end login functionality

// Logout functionality...

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	fmt.Println("User logout successfuly!")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout handler runing...")
	clearSession(w)
	http.Redirect(w, r, "/index", 302)
}

// end logout functionality

//Dashboard functionality...

func getUserName(request *http.Request) (userName string) {
	fmt.Println("fetch cookies..")
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}

func Dashboardhandler(w http.ResponseWriter, r *http.Request) {
	email := getUserName(r)
	fmt.Println("Dashboard handler running..", email)
	p1 := person{
		email,
	}
	fmt.Println(p1)
	if (p1.First != "") || (person{} != p1) {
		fmt.Println(p1)

		db := dbc.DbConn()
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
			var id, retailer_id int
			var name, brand, description, created_date string
			var price float32
			var status bool
			err = selDB.Scan(&id, &retailer_id, &name, &brand, &price, &description, &created_date, &status)
			if err != nil {
				panic(err.Error())
			}
			pro.Id = id
			pro.Name = name
			pro.Brand = brand
			pro.Price = price
			pro.Description = description

			t, _ := time.Parse("2006-01-02", created_date)
			pro.Created_date = t
			pro.Status = status
			res = append(res, pro)
		}
		data := map[string]interface{}{
			"p1":  p1,
			"pro": pro,
		}
		fmt.Println(data)
		tmpl.ExecuteTemplate(w, "Productlist.html", data)

	} else {
		fmt.Println("not found")
		tmpl.ExecuteTemplate(w, "Index.html", nil)

	}
}

// end Dashboard functionality
