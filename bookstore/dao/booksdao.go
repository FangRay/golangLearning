package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

func ShowBooks() ([]*model.Book, error) {

	var books []*model.Book
	sqlStr := "select id,title,author,price,sales,stock,imgpath from books"
	rows, err := utils.DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.Imgpath)
		books = append(books, book)
	}

	//fmt.Println(books)
	return books, nil

}

func DeleteBook(bookid string) error {
	sqlStr := "delete from books where id=?"
	_, err := utils.DB.Exec(sqlStr, bookid)
	if err != nil {
		return err
	}
	return nil
}

func FindBook(bookid string) (*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,imgpath from books where id = ?"
	row := utils.DB.QueryRow(sqlStr, bookid)

	book := &model.Book{}
	row.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.Imgpath)
	return book, nil
}

func UpdateBook(b *model.Book) error {
	if b.Id > 0 {
		sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
		utils.DB.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.Id)
		return nil
	}
	sqlStr := "insert into books(title,author,price,sales,stock,imgpath)values(?,?,?,?,?,?)"
	utils.DB.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.Imgpath)
	return nil

}

func PageBook(pageno int64) (*model.Paginator, error) {
	//查出总数量，插入paginator结构体
	//fmt.Println("开始分页")
	sqlStr := "select count(*) from books"
	row := utils.DB.QueryRow(sqlStr)
	var pagesum int64
	row.Scan(&pagesum)
	//fmt.Printf("总数是%d", pagesum)

	//制造结构体里的 book结构体切片
	var bookslice []*model.Book
	sqlStr = "select id,title,author,price,sales,stock,imgpath from books limit ?,?"
	rows, _ := utils.DB.Query(sqlStr, (pageno-1)*4, 4)
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.Imgpath)
		bookslice = append(bookslice, book)
	}
	paginator := &model.Paginator{
		Bookslice: bookslice,
		Pageno:    pageno,
		Pagesum:   pagesum,
	}
	return paginator, nil

}
