// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// ArticleData is the golang structure for table article_data.
type ArticleData struct {
	Id         uint64      `orm:"id,primary"  json:"id"`          //
	ArticleId  uint64      `orm:"article_id"  json:"article_id"`  // 文章ID
	SubTitle   string      `orm:"sub_title"   json:"sub_title"`   // 副标题
	Answer     string      `orm:"answer"      json:"answer"`      // 答案
	Body       string      `orm:"body"        json:"body"`        // 主题内容
	Analysis   string      `orm:"analysis"    json:"analysis"`    // 解析内容
	CreateTime *gtime.Time `orm:"create_time" json:"create_time"` // 创建时间
	ModifyTime *gtime.Time `orm:"modify_time" json:"modify_time"` // 更新时间
}
