package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"
	dbc "video/db"
)

type User struct {
	Id         int
	First_Name string
	Last_Name  string
	Email      string
	Gender     string
	DOB        time.Time
}

func Greet() string {
	return "Server started..."
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index hendler running...")
	tmpl.ExecuteTemplate(w, "Index.html", nil)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signup handler running...")
	db := dbc.DbConn()

	if r.Method == "POST" {
		first_name := r.FormValue("first_name")
		last_name := r.FormValue("last_name")
		email := r.FormValue("email")
		dob := r.FormValue("dob")
		gender := r.FormValue("gender")
		status := "1"
		insForm, err := db.Prepare("INSERT INTO User(first_name, last_name,email,dob,gender,status) VALUES(?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(first_name, last_name, email, dob, gender, status)
		log.Println("INSERT: First_Name: " + first_name + " | Last_Name: " + last_name + " | Email: " + email + " | DOB:" + dob + " | Gender: " + gender + " | Status: " + status)
	}
	defer db.Close()
	http.Redirect(w, r, "/Index", 301)

}
