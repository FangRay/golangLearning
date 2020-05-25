package parse

import (
	"pachong/engine"
	"pachong/model"
	"regexp"
)

func ParseMain(bytes []byte) engine.ParseResult {
	var weylin string = `/(.[0-9]+/)">([^<]+)</a>`
	compile := regexp.MustCompile(weylin)

	//allString := compile.FindAll(result,-1)
	matches := compile.FindAllSubmatch(bytes, -1)
	result := engine.ParseResult{}

	for _, v := range matches {
		url := `https://www.biquge.com.cn/book/` + string(v[1])
		result.Items = append(result.Items, string(v[2]))
		// 老师这个方法好，直接用append里生产结构体
		result.Requests = append(result.Requests, engine.Requests{
			Url:       url,
			ParseFunc: SecondPage,
		})
		//把书信息提出来的requests
		bookname := string(v[2])
		result.Requests = append(result.Requests, engine.Requests{
			Url: url,
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return BookInfo(bytes, bookname)
			},
		})
	}

	/*	for _, v := range matches {
		url := `https://www.biquge.com.cn/book/` + string(v[1])
		bookname := string(v[2])
		result.Requests = append(result.Requests, engine.Requests{
			Url: url,   //因为 bookinfo有2参数，但是Parsefun只有1参数，所以建立一个匿名函数，闭包
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return BookInfo(bytes, bookname) //for里的函数，参数一定要复制一下
			},
		})
	}*/

	return result
}

func BookInfo(bytes []byte, pname string) engine.ParseResult {
	newbook := model.Bookinfo{}
	pic := `og:image" content="(https://www.biquge.com.cn/files/article/image/[^.]+.jpg)"/>`
	piceg := regexp.MustCompile(pic)
	booktype := `og:novel:category" content="([^"]+)"/>`
	booktypege := regexp.MustCompile(booktype)
	author := `og:novel:author" content="([^"]+)"/>`
	authoreg := regexp.MustCompile(author)
	name := `og:novel:book_name" content="([^"]+)"/>`
	nameeg := regexp.MustCompile(name)
	url := `og:url" content="(https://www.biquge.com.cn/[^"]+)"/>`
	urleg := regexp.MustCompile(url)
	newbook.Url = getstring(bytes, urleg)
	newbook.Author = getstring(bytes, authoreg)
	newbook.Name = getstring(bytes, nameeg)
	newbook.Pic = getstring(bytes, piceg)
	newbook.Type = getstring(bytes, booktypege)
	newbook.Passname = pname
	result := engine.ParseResult{}
	result.Items = append(result.Items, newbook)

	/*也可以用这种方式 ：   result:=engine.ParseResult{Items: []interface{}{newbook},}*/
	return result
}

func SecondPage(bytes []byte) engine.ParseResult {
	var reg string = `/book/([\d]+/)([\d]+.html)[^>]+>([^<]+)</a>`
	compile := regexp.MustCompile(reg)
	submatch := compile.FindAllSubmatch(bytes, -1)
	result := engine.ParseResult{}
	var count int
	for _, v := range submatch {
		result.Items = append(result.Items, string(v[3]))
		url := `https://www.biquge.com.cn/book/` + string(v[1]) + string(v[2])
		//log.Printf("Second Page ParseLink %s :%s", url, string(v[3]))
		result.Requests = append(result.Requests, engine.Requests{
			Url:       url,
			ParseFunc: NilParse,
		})

		count++
		if count > 10 {
			break
		}
	}
	//把页面所有其它书的连接都提取出来 ，重复mainpage的步骤而已
	var weylin string = `/([a-z0-9]+/)">([^<]+)</a>`
	compile2 := regexp.MustCompile(weylin)

	//allString := compile.FindAll(result,-1)
	matches2 := compile2.FindAllSubmatch(bytes, -1)
	for _, v := range matches2 {
		url2 := `https://www.biquge.com.cn/book/` + string(v[1])
		result.Requests = append(result.Requests, engine.Requests{
			Url:       url2,
			ParseFunc: SecondPage,
		})
		result.Items = append(result.Items, string(v[2]))
	}

	return result

}

func getstring(bytes []byte, re *regexp.Regexp) string {
	result := re.FindSubmatch(bytes)
	if result != nil {
		return string(result[1])
	}
	return ""
}

func NilParse([]byte) engine.ParseResult {
	return engine.ParseResult{}

}
