// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// ThemeMap is the golang structure for table theme_map.
type ThemeMap struct {
	Id         uint64      `orm:"id,primary"  json:"id"`          //
	ThemeId    uint64      `orm:"theme_id"    json:"theme_id"`    // 标签ID
	ArticleId  uint64      `orm:"article_id"  json:"article_id"`  // 内容ID
	AppId      uint        `orm:"app_id"      json:"app_id"`      // 应用ID
	CreateTime *gtime.Time `orm:"create_time" json:"create_time"` // 创建时间
	ModifyTime *gtime.Time `orm:"modify_time" json:"modify_time"` // 更新时间
}