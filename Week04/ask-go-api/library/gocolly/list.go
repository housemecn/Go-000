package gocolly

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"golang.org/x/net/http2"

	"git.code.tencent.com/houseme/ask-go-api/app/dao"
	"git.code.tencent.com/houseme/ask-go-api/app/model"
	"git.code.tencent.com/houseme/ask-go-api/library/redis"
)

// @Project: ask-go-api
// @Author: houseme
// @Description:
// @File: list
// @Version: 1.0.0
// @Date: 2020/12/5 16:50
// @Package gocolly

// SelectCategory Select Category
func SelectCategory(conn *gredis.Conn) {
	list, _ := dao.Category.Where("id>31").Order("id desc").All()
	var page uint = 1
	var total = 0
	for k, v := range list {
		if v.TotalPage < 2 {
			continue
		}
		fmt.Println("category:", v, " k:", k)
		page = v.CurrentPage
		for ; page <= v.TotalPage; page++ {
			ListColly(conn, v.Id, page, v.SourceUrl)
			//time.Sleep(3 * time.Second)
			total = total + 1
			fmt.Println("time:", time.Now(), "total:", total)
			if total%10 == 3 {
				time.Sleep(1 * time.Second)
			}
		}
		fmt.Println("time:", time.Now(), "total:", total)
	}
}

// ListColly List Colly
func ListColly(conn *gredis.Conn, id, page uint, URL string) {
	var start = gtime.Now().UnixNano()
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
	c.OnHTML(".listleft", func(e *colly.HTMLElement) {
		e.ForEach(".p30", func(i int, element *colly.HTMLElement) {
			fmt.Println("article")
			href, _ := element.DOM.Find("a[href]").Attr("href")
			sourceURL := strings.TrimSpace(BaseURL + href)
			mds := fmt.Sprintf("%x", md5.Sum([]byte(sourceURL)))
			fmt.Println("URL:", sourceURL, "mds:", mds, " href:", href)
			v, _ := conn.DoVar("GET", redis.ArticleTitle+mds)
			if v.IsNil() || v.IsEmpty() {
				var article *model.Article
				article, err = dao.Article.FindOne("md5_value=?", mds)
				if article == nil {
					article = new(model.Article)
					article.Cid = id
					article.Description = element.DOM.Find(".dnone").Text()
					article.Title = element.DOM.Find("h2").Text()
					article.SourceUrl = sourceURL
					article.SortNum = uint(i)
					article.Md5Value = mds
					article.State = 100
					result, err := dao.Article.Data(article).Unscoped().OmitEmpty().Insert()
					fmt.Println("dao.Article insert result:", result, " error:", err)
				}
				conn.Do("SET", redis.ArticleTitle+mds, gtime.Now().Unix())
				v, _ := conn.DoVar("GET", redis.ThemeTitle+mds)
				fmt.Println("redis article key:", redis.ArticleTitle+mds, " value:", v.String())
			}
			fmt.Println("url:", href, " title: ", element.DOM.Find("h2").Text(), "ID:", i, "err:", err)
		})
	})

	c.OnHTML("#relateDiv", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
			fmt.Println("theme")
			sourceURL := strings.TrimSpace(BaseURL + element.Attr("href"))
			mds := fmt.Sprintf("%x", md5.Sum([]byte(sourceURL)))
			v, _ := conn.DoVar("GET", redis.ThemeTitle+mds)
			if v.IsNil() || v.IsEmpty() {
				var theme *model.Theme
				theme, err = dao.Theme.FindOne("md5_value=?", mds)
				if theme == nil {
					theme = new(model.Theme)
					theme.Name = element.Text
					theme.Title = element.Text
					theme.Dir = element.Attr("href")
					theme.SourceUrl = sourceURL
					theme.SortNum = uint(i)
					theme.Md5Value = mds
					theme.State = 100
					dao.Theme.Data(theme).Unscoped().OmitEmpty().Insert()
				}
				conn.Do("SET", redis.ThemeTitle+mds, gtime.Now().Unix())
				v, _ := conn.DoVar("GET", redis.ThemeTitle+mds)
				fmt.Println("redis theme key:", redis.ThemeTitle+mds, " value:", v.String())
			}
			fmt.Println("url:", element.Attr("href"), " title: ", element.Text, "ID:", i, "err:", err)
		})
	})

	c.OnHTML(".am-pagination", func(e *colly.HTMLElement) {
		var pageContent = e.DOM.Find(".pageRemark")
		totalPage, _ := strconv.Atoi(pageContent.Find("b").First().Text())
		total, _ := strconv.Atoi(pageContent.Find("b").Last().Text())
		dao.Category.Data(g.Map{
			"total":        total,
			"total_page":   totalPage,
			"current_page": page,
			"modify_time":  gtime.Now(),
		}).Where("id=?", id).Unscoped().Update()
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	//c.Visit("https://www.asklib.com/")
	c.OnScraped(func(r *colly.Response) {
		//fmt.Println("Body: ", string(r.Body[:]))
		end := gtime.Now().UnixNano()
		fmt.Println("end time:", end, " all time:", end-start)

	})
	if page > 1 {
		var buffer bytes.Buffer
		buffer.WriteString(strings.ReplaceAll(URL, ".html", "/p"))
		buffer.WriteString(strconv.Itoa(int(page)))
		buffer.WriteString(".html")
		c.Visit(buffer.String())
	} else {
		c.Visit(URL)
	}
}
