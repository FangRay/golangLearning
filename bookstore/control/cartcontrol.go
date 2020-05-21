package control

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

//要检查一下，是用postFormValue, 还是formvalue
func AddToCart(w http.ResponseWriter, r *http.Request) {
	//fmt.Print("hello add to cart")
	temp, _ := r.Cookie("user")
	if temp == nil {
		// t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		// t.Execute(w, "请登陆")   , not sucess, because not Jquery support
		w.Write([]byte("Login Please"))
	} else {
		sess := temp.Value
		//fmt.Printf("第83行sess的值是%v\n", sess)
		flag, _, userid := model.IsLogin(sess)
		if !flag {
			//fmt.Printf("第86行判断登陆的结果是flag%v\n", flag)
			//在jquery $post里无法跳转动作
			// t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			// fmt.Print(t)
			// t.Execute(w, "")
			w.Write([]byte("Login Please"))
		} else {
			tempstr := r.PostFormValue("bookid")
			tempint64, _ := strconv.ParseInt(tempstr, 0, 0)
			goodid := int(tempint64)
			goodname := r.PostFormValue("bookname")
			tempstr = r.PostFormValue("bookprice")
			//fmt.Println(tempstr)
			goodprice, _ := strconv.ParseFloat(tempstr, 64)
			//fmt.Println(goodprice)
			//try to get the right AJAX info
			//fmt.Printf("第93行书的获取信息为%v,%v,%f", goodid, goodname, goodprice)
			var goodcount int = 1
			newcart := &model.Cart{
				Userid:    userid,
				Goodid:    goodid,
				Goodname:  goodname,
				Goodprice: float32(goodprice),
				Count:     goodcount,
			}
			//print the new cart struct, success!
			fmt.Println(newcart)
			newcart.AddCart()
			w.Write([]byte("You just add【 " + goodname + " 】to you cart, thanks!"))
			//打印一下此人购物车的总项
			//fmt.Print(newcart.CartTotal())
		}

	}

}

func ShowCart(w http.ResponseWriter, r *http.Request) {
	//
	//fmt.Println("show me memem")
	Cookie, _ := r.Cookie("user")
	if Cookie == nil {
		Mainhandle(w, r)
	} else {
		sess := Cookie.Value
		flag, username, userid := model.IsLogin(sess)
		if !flag {
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "请登陆")
		} else {
			newcart := &model.Cart{
				Userid: userid,
			}
			//比较笨的方法去取得用户名等信息，因为如果传过去切片的话，只能对里面的内容进行函数操作
			sumpiece, sumtotal := newcart.CartTotal()
			cartslice := newcart.GetCart()
			ShowCart := &model.ShowCart{
				Cartslice: cartslice,
				Username:  username,
				Sumpiece:  sumpiece,
				Sumprice:  sumtotal,
				Userid:    userid,
			}
			t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
			t.Execute(w, ShowCart)
		}

	}
}

func DeleteGood(w http.ResponseWriter, r *http.Request) {
	temp := r.FormValue("goodid")
	//
	//fmt.Printf("93:goodidvalue:%v\n", temp)
	temp64, _ := strconv.ParseInt(temp, 0, 0)
	goodid := int(temp64)
	temp = r.FormValue("userid")
	temp64, _ = strconv.ParseInt(temp, 0, 0)
	userid := int(temp64)
	//
	//fmt.Printf("98:goodid%v,userid%v\n", goodid, userid)
	model.DeleteGood(goodid, userid)
	ShowCart(w, r)
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	temp := r.FormValue("userid")
	temp64, _ := strconv.ParseInt(temp, 0, 0)
	userid := int(temp64)
	sql := "delete from carts where userid =?"
	utils.DB.Exec(sql, userid)
	ShowCart(w, r)
}

func ModifyCart(w http.ResponseWriter, r *http.Request) {
	temp := r.FormValue("userid")
	temp64, _ := strconv.ParseInt(temp, 0, 0)
	userid := int(temp64)

	temp = r.FormValue("goodid")
	temp64, _ = strconv.ParseInt(temp, 0, 0)
	goodid := int(temp64)

	temp = r.FormValue("goodcount")
	temp64, _ = strconv.ParseInt(temp, 0, 0)
	goodcount := int(temp64)

	//test if input float but not int value, success!
	fmt.Printf("133:paravaluesare %v,%v,%v \n", userid, goodid, goodcount)

	if goodcount < 1 || goodcount > 100 {
		w.Write([]byte("wrongnumber"))

	} else {
		sql := "update carts set count=? where userid=? and goodid=?"
		_, err := utils.DB.Exec(sql, goodcount, userid, goodid)
		if err != nil {
			w.Write([]byte("WTF"))
		} else {
			w.Write([]byte("right"))
		}
	}

}
