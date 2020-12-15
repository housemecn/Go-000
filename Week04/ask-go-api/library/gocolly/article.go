package gocolly

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
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
// @File: article
// @Version: 1.0.0
// @Date: 2020/12/6 00:54
// @Package gocolly

// SelectArticle Select Category
func SelectArticle(conn *gredis.Conn) {
	//list, _ := dao.Article.Where("id<100000 and id >1").Order("id asc").All()
	//db := g.DB()
	//rows, err := db.Query("SELECT * FROM article a LEFT JOIN article_data d ON  a.id = d.article_id  WHERE a.id<5031 AND d.id IS NULL")
	var total = 0
	var i uint64
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i = 1; i < 960727; i++ {
		fmt.Println("Article ID:", i)
		if i < 82542 {
			articleData, _ := dao.ArticleData.Where("article_id=?", i).FindOne()
			if articleData != nil {
				continue
			}
		}
		article, _ := dao.Article.Where("id=?", i).FindOne()
		if article == nil {
			continue
		}
		ArticleColly(conn, article)
		total = total + 1
		fmt.Println("Article time:", time.Now(), "total:", total)
		num := time.Duration(r.Intn(5)) * time.Second
		if total%10 == 2 {
			time.Sleep(num)
		}
		fmt.Println("Article time:", time.Now(), "total:", total, " num:", num)
	}
}

// ArticleColly List Colly
func ArticleColly(conn *gredis.Conn, article *model.Article) {
	c := colly.NewCollector(
		// 限定域名
		colly.AllowedDomains("www.asklib.com", "asklib.com", "m.asklib.com"),
		// 最大深度
		colly.MaxDepth(1),
		colly.Async(true),
	)
	c.WithTransport(&http2.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	})
	c.SetRequestTimeout(50 * time.Second)
	extensions.RandomUserAgent(c)
	c.Async = true

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
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	// Find and visit all links
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//fmt.Println("url:", e.Attr("href"), " title: ", e.Text)
	//e.Request.Visit(e.Attr("href"))
	//})
	var err error
	/*c.OnHTML(".listleft", func(e *colly.HTMLElement) {
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
	})*/

	c.OnHTML("#relateDiv", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
			fmt.Println("theme")
			sourceURL := strings.TrimSpace(BaseURL + element.Attr("href"))
			mds := fmt.Sprintf("%x", md5.Sum([]byte(sourceURL)))
			fmt.Println("sourceURL:", sourceURL, "mds:", mds)
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
				v, err := conn.DoVar("GET", redis.ThemeTitle+mds)
				fmt.Println("redis theme key:", redis.ThemeTitle+mds, " value:", v.String(), " error:", err)
			}
			fmt.Println("url:", element.Attr("href"), " title: ", element.Text, "ID:", i, "err:", err)
		})
	})

	//c.OnHTML(".am-pagination", func(e *colly.HTMLElement) {
	//    var pageContent = e.DOM.Find(".pageRemark")
	//    totalPage, _ := strconv.Atoi(pageContent.Find("b").First().Text())
	//    total, _ := strconv.Atoi(pageContent.Find("b").Last().Text())
	//    dao.Category.Data(g.Map{
	//        "total":        total,
	//        "total_page":   totalPage,
	//        "current_page": page,
	//        "modify_time":  gtime.Now(),
	//    }).Where("id=?", id).Unscoped().Update()
	//})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		var doc *goquery.Document
		doc, err = goquery.NewDocumentFromReader(bytes.NewBuffer(r.Body[:]))
		keywords, _ := doc.Find("meta[name='keywords']").Attr("content")
		dao.Article.Data(g.Map{
			"keywords":    keywords,
			"modify_time": gtime.Now(),
		}).Where("id=?", article.Id).Update()
		var articleData = new(model.ArticleData)
		articleData.ArticleId = article.Id
		articleData.Body, err = doc.Find(".essaytitle").Html()
		articleData.Answer, err = doc.Find(".listbg").Html()
		if len(articleData.Answer) < 1 {
			articleData.Answer = ""
		}
		articleData.Analysis, err = doc.Find("#commentDiv").Html()
		var validLi = regexp.MustCompile(`<script(.[\s\S]*?)</script>`)
		as := validLi.FindAllString(articleData.Analysis, -1)
		if len(as) > 0 {
			for di, a := range as {
				articleData.Analysis = strings.ReplaceAll(articleData.Analysis, a, "")
				fmt.Println("di: ", di, " a: ", a)
			}
		}
		var validIns = regexp.MustCompile(`<ins(.[\s\S]*?)</ins>`)
		as = validIns.FindAllString(articleData.Analysis, -1)
		if len(as) > 0 {
			for di, a := range as {
				articleData.Analysis = strings.ReplaceAll(articleData.Analysis, a, "")
				fmt.Println("di: ", di, " a: ", a)
			}
		}
		dao.ArticleData.Data(articleData).Unscoped().OmitEmpty().Insert()
	})

	/*c.Limit(&colly.LimitRule{
		DomainGlob:  "*asklib.*",
		Parallelism: 18,
		RandomDelay: 5 * time.Second,
	})*/

	c.Visit(article.SourceUrl)
	c.Wait()
}
