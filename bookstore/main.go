package main

import (
	"bookstore/control"
	"net/http"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/", control.Mainhandle)
	http.HandleFunc("/login", control.Login)
	http.HandleFunc("/logout", control.LogOut)
	http.HandleFunc("/regist", control.Regist)
	http.HandleFunc("/checkusername", control.CheckUserName)
	http.HandleFunc("/managebooks", control.ManageBooks)
	http.HandleFunc("/deletebook", control.DeleteBook)
	http.HandleFunc("/editbook", control.EditBook)
	http.HandleFunc("/updatebook", control.UpdateBook)
	http.HandleFunc("/addcart", control.AddToCart)
	http.HandleFunc("/showcart", control.ShowCart)
	http.HandleFunc("/deletegood", control.DeleteGood)
	http.HandleFunc("/deletecart", control.DeleteCart)
	http.HandleFunc("/modifycart", control.ModifyCart)

	http.ListenAndServe(":8080", nil)

}
