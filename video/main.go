package main

import (
	"fmt"
	"net/http"
	handler "video/handler"
)

func main() {
	fmt.Println(handler.Greet())

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/index", handler.IndexHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/logout", handler.LogoutHandler)
	http.HandleFunc("/insert", handler.SignupHandler)
	http.HandleFunc("/dashboard", handler.Dashboardhandler)
	http.HandleFunc("/showproduct", handler.ShowproductHandler)
	http.HandleFunc("/addproduct", handler.AddproductHandler)

	http.ListenAndServe(":8000", nil)

}
