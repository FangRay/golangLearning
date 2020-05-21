package model

import (
	"bookstore/utils"
)

type Cart struct {
	Userid    int
	Goodid    int
	Goodname  string
	Goodprice float32
	Count     int
	Username  string
}

type ShowCart struct {
	Cartslice []*Cart
	Username  string
	Sumpiece  int
	Sumprice  float32
	Userid    int
}

//sql代码
// create table carts(
// 	userid int not null,
// 	goodid int not null,
// 	goodname varchar(100) not null,
// 	goodprice double(11,2) not null,
// 	foreign key(goodid) references books(id),
// 	foreign key (userid) references users(id)
// 	);
func (c *Cart) AddCart() {
	var count int
	strsql := "select count from carts where goodid =? and userid =?"
	row := utils.DB.QueryRow(strsql, c.Goodid, c.Userid)
	row.Scan(&count)
	//fmt.Printf("addcart38行count数是%", count)
	if count > 0 {
		strsql = "update carts set count =? where goodid=? and userid=?"
		utils.DB.Exec(strsql, count+1, c.Goodid, c.Userid)
	} else {
		strsql = "insert into carts(userid,goodid,goodname,goodprice,count) values(?,?,?,?,?)"
		//mysql注意编码问题，插中文要GBK
		//fmt.Printf("Addcart:45:insert values are %v,%v,%v,%v,%v\n", c.Userid, c.Goodid, c.Goodname, c.Goodprice, c.Count)
		utils.DB.Exec(strsql, c.Userid, c.Goodid, c.Goodname, c.Goodprice, c.Count)

	}

}

func (c *Cart) GetUserId() int {
	return c.Userid
}

func (c *Cart) GetAmount() float32 {
	return float32(c.Count) * c.Goodprice
}

//GetCart通过结构体里的用户ID，得到所有这个用户ID的购物明细切片
func (c *Cart) GetCart() []*Cart {
	cartslice := []*Cart{}

	strsql := "select goodid,goodname,goodprice,count from carts where userid=?"
	rows, _ := utils.DB.Query(strsql, c.Userid)
	for rows.Next() {
		cart := &Cart{
			Userid: c.Userid,
			//如果是下面getTotal方法里并没有加入Username字段的话，是空也没有关系
			Username: c.Username,
		}
		rows.Scan(&cart.Goodid, &cart.Goodname, &cart.Goodprice, &cart.Count)
		cartslice = append(cartslice, cart)
	}
	return cartslice

}

//到时候试一下要不要分成两个方法，还是可以直接用.totalxxx来套模版
func (c *Cart) CartTotal() (totalpiece int, totalvalue float32) {
	cartslice := c.GetCart()

	for _, v := range cartslice {
		//fmt.Print(v)
		totalpiece += v.Count
		//64位精度太高会错
		totalvalue += v.GetAmount()
		//fmt.Printf("t:%v ,totova:%v\n", t, totalvalue)
	}
	return
}

func DeleteGood(goodid int, userid int) error {
	strsql := "delete from carts where goodid=? and userid=?"
	_, err := utils.DB.Exec(strsql, goodid, userid)
	if err != nil {
		return err
	}
	return nil
}
