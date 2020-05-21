package control

import (
	"bookstore/dao"
	"bookstore/model"
	"net/http"
	"strconv"
	"text/template"
)

func Mainhandle(w http.ResponseWriter, r *http.Request) {

	rpageno := r.FormValue("pageno")
	if rpageno == "" {
		rpageno = "1"
	}
	//fmt.Println(rpageno)

	pageno, _ := strconv.ParseInt(rpageno, 10, 64)
	paginator, _ := dao.PageBook(pageno)
	//向paginator里插入session结构体
	temp, _ := r.Cookie("user")
	if temp != nil {
		var sess string
		sess = temp.Value
		flag, username, userid := model.IsLogin(sess)
		if flag {
			paginator.Islogin = true
			paginator.Username = username
			cart := &model.Cart{
				Userid: userid,
			}
			paginator.Cartpieces, _ = cart.CartTotal()

		}

	}

	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, paginator)
}

func ManageBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.ShowBooks()
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, books)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookid := r.FormValue("bookid")
	dao.DeleteBook(bookid)
	ManageBooks(w, r)
}

func EditBook(w http.ResponseWriter, r *http.Request) {
	bookid := r.FormValue("bookid")
	book, _ := dao.FindBook(bookid)
	t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
	t.Execute(w, book)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	iid, _ := strconv.ParseInt(id, 0, 0)
	fprice, _ := strconv.ParseFloat(price, 64)
	isale, _ := strconv.ParseInt(sales, 10, 0)
	istock, _ := strconv.ParseInt(stock, 10, 0)
	book := &model.Book{
		Id:     int(iid),
		Title:  title,
		Author: author,
		Price:  float32(fprice),
		Sales:  int(isale),
		Stock:  int(istock),
	}
	dao.UpdateBook(book)
	ManageBooks(w, r)

}
