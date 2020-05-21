package model

type Paginator struct {
	Bookslice []*Book
	Pageno    int64
	Pagesum   int64
	//Sess      *Session
	Islogin  bool
	Username string
	//购物车信息
	Cartpieces int
}

func (p *Paginator) PageTotal() int64 {
	if p.Pagesum%4 == 0 {
		return p.Pagesum / 4
	}
	return p.Pagesum/4 + 1

}

func (p *Paginator) FirstPage() bool {
	return p.Pageno != 1
}

func (p *Paginator) EndPage() bool {
	return p.Pageno != p.PageTotal()
}

func (p *Paginator) LastPage() int64 {
	return p.Pageno - 1
}

func (p *Paginator) NextPage() int64 {
	return p.Pageno + 1
}
