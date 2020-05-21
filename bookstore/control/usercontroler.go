package control

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {

	tempstr, _ := r.Cookie("user")
	fmt.Printf("15行temp值%v", tempstr)
	if tempstr != nil {
		sess := tempstr.Value
		flag, _, _ := model.IsLogin(sess)
		if flag {
			Mainhandle(w, r)
		}
	}

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, _ := dao.CheckUser(username, password)
	if user.ID > 0 {
		uid := utils.CreatUid()
		fmt.Print(uid)
		session := &model.Session{
			Sess:     uid,
			Username: username,
			Userid:   user.ID,
		}

		session.Add()
		//向浏览器插入cookie
		cookie := &http.Cookie{
			Name:     "user",
			Value:    uid,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)

		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, session)
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "不正确")
	}

}

func LogOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1我要清除session啦")
	temp, _ := r.Cookie("user")
	cookie := temp.Value
	fmt.Println(cookie)
	newsession := &model.Session{
		Sess: cookie,
	}
	fmt.Println("我要清除session啦")
	fmt.Println(newsession)
	newsession.Quit()
	Mainhandle(w, r)

}

func Regist(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	if dao.RegistUser(username, password, email) != nil {
		fmt.Println("something wrong")
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}

}

func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		w.Write([]byte("用户名已经存在"))
	} else {
		w.Write([]byte("<font style='color:green'>用户名可以注册</font>"))
	}

}
