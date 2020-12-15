package gocolly

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"golang.org/x/net/http2"

	"git.code.tencent.com/houseme/ask-go-api/app/dao"
	"git.code.tencent.com/houseme/ask-go-api/app/model"
)

// @Project: ask-go-api
// @Author: houseme
// @Description:
// @File: gocolly
// @Version: 1.0.0
// @Date: 2020/12/4 00:46
// @Package gocolly

// Colly go colly
func Colly() {
	c := colly.NewCollector(
		// 限定域名
		colly.AllowedDomains("www.asklib.com", "asklib.com", "m.asklib.com"),
		// 最大深度
		colly.MaxDepth(1),
	)
	c.WithTransport(&http2.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})
	c.SetRequestTimeout(30 * time.Second)
	extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("user-agent", UserAgent)
		r.Headers.Set("accept", Accept)
		r.Headers.Set("accept-Encoding", AcceptEncoding)
		r.Headers.Set("authority", Authority)
		r.Headers.Set("cookie", Cookie)
		r.Headers.Set("accept-language", AcceptLanguage)
		r.Headers.Set("referer", Referer)
		r.Headers.Set("sec-fetch-dest", SecFetchDest)
		r.Headers.Set("sec-fetch-mode", SecFetchMode)
		r.Headers.Set("sec-fetch-site", SecFetchSite)
		r.Headers.Set("sec-fetch-user", SecFetchUser)
		r.Headers.Set("upgrade-insecure-requests", UpgradeInsecureRequests)
		r.Headers.Set("Connection", Connection)
		fmt.Println("Visiting", r.URL)
	})

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//fmt.Println("url:", e.Attr("href"), " title: ", e.Text)
		//e.Request.Visit(e.Attr("href"))
	})
	var err error
	c.OnHTML(".boxcon", func(e *colly.HTMLElement) {
		//d :=e.DOM.Find(".detail").Eq(0)
		e.ForEach(".detail", func(id int, d *colly.HTMLElement) {
			d.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
				sourceURL := BaseURL + element.Attr("href")
				mds := fmt.Sprintf("%x", md5.Sum([]byte(sourceURL)))
				if id == 0 {
					fmt.Println("category")
					var cate *model.Category
					cate, err = dao.Category.FindOne("md5_value=?", mds)
					if cate == nil {
						cate = new(model.Category)
						cate.Name = element.Text
						cate.Title = element.Text
						cate.Title = element.Text
						cate.Dir = element.Attr("href")
						cate.SourceUrl = sourceURL
						cate.SortNum = uint(i)
						cate.Md5Value = mds
						cate.State = 100
						//cate.CreateTime = gtime.Now()
						//cate.ModifyTime = gtime.Now()
						dao.Category.Data(cate).Unscoped().OmitEmpty().Insert()
					}
				} else {
					fmt.Println("theme")
					var theme *model.Theme
					theme, err = dao.Theme.FindOne("md5_value=?", mds)
					if theme == nil {
						theme = new(model.Theme)
						theme.Name = element.Text
						theme.Title = element.Text
						theme.Title = element.Text
						theme.Dir = element.Attr("href")
						theme.SourceUrl = sourceURL
						theme.SortNum = uint(i)
						theme.Md5Value = mds
						//theme.CreateTime = gtime.Now()
						//theme.ModifyTime = gtime.Now()
						theme.State = 100
						dao.Theme.Data(theme).Unscoped().OmitEmpty().Insert()
					}
				}
				fmt.Println("url:", element.Attr("href"), " title: ", element.Text, "ID:", i, "err:", err)
			})
		})
	})

	c.OnHTML(".searchtype", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
			fmt.Println("category")
			sourceURL := BaseURL + element.Attr("href")
			mds := fmt.Sprintf("%x", md5.Sum([]byte(sourceURL)))
			var cate *model.Category
			cate, err = dao.Category.FindOne("md5_value=?", mds)
			if cate == nil {
				cate = new(model.Category)
				cate.Name = element.Text
				cate.Title = element.Text
				cate.Title = element.Text
				cate.Dir = element.Attr("href")
				cate.SourceUrl = sourceURL
				cate.SortNum = uint(i)
				cate.Md5Value = mds
				cate.State = 100
				dao.Category.Data(cate).Unscoped().OmitEmpty().Insert()
			}
			fmt.Println("url:", element.Attr("href"), " title: ", element.Text, "ID:", i, "err:", err)
		})
	})

	c.OnHTML(".mb20 >ul", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
			fmt.Println("category")
			if element.Text != "网站首页" {
				sourceURL := BaseURL + element.Attr("href")
				mds := fmt.Sprintf("%x", md5.Sum([]byte(sourceURL)))
				var cate *model.Category
				cate, err = dao.Category.FindOne("md5_value=?", mds)
				if cate == nil {
					cate = new(model.Category)
					cate.Name = element.Text
					cate.Title = element.Text
					cate.Title = element.Text
					cate.Dir = element.Attr("href")
					cate.SourceUrl = sourceURL
					cate.SortNum = uint(i)
					cate.Md5Value = mds
					cate.State = 100
					dao.Category.Data(cate).Unscoped().OmitEmpty().Insert()
				}
			}
			fmt.Println("url:", element.Attr("href"), " title: ", element.Text, "ID:", i, "err:", err)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	//c.Visit("https://www.asklib.com/")
	c.Visit("https://www.asklib.com/kaoyan.html")
}
