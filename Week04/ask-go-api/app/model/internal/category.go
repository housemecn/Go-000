// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// Category is the golang structure for table category.
type Category struct {
	Id          uint        `orm:"id,primary"       json:"id"`           // 分类id
	Name        string      `orm:"name"             json:"name"`         // 分类名称
	SubName     string      `orm:"sub_name"         json:"sub_name"`     // 子名称
	SortNum     uint        `orm:"sort_num"         json:"sort_num"`     // 排序值
	Title       string      `orm:"title"            json:"title"`        // 标题
	Keywords    string      `orm:"keywords"         json:"keywords"`     // 关键字
	Description string      `orm:"description"      json:"description"`  // 描述
	Dir         string      `orm:"dir"              json:"dir"`          // 目录
	Count       uint        `orm:"count"            json:"count"`        // 数量
	LabelOne    string      `orm:"label_one"        json:"label_one"`    // 标签1 品牌页的顶部标题
	LabelTwo    string      `orm:"label_two"        json:"label_two"`    // 标签2 品牌中间标题
	Remark      string      `orm:"remark"           json:"remark"`       // 备注
	Md5Value    string      `orm:"md5_value,unique" json:"md_5_value"`   // url md5值
	SourceUrl   string      `orm:"source_url"       json:"source_url"`   // 原始URL
	CurrentPage uint        `orm:"current_page"     json:"current_page"` // 分类当前页
	Total       uint        `orm:"total"            json:"total"`        // 总内容数
	TotalPage   uint        `orm:"total_page"       json:"total_page"`   // 分类总页数
	State       uint        `orm:"state"            json:"state"`        // 状态 100默认正常 90代审核 120禁用
	CreateTime  *gtime.Time `orm:"create_time"      json:"create_time"`  // 创建时间
	ModifyTime  *gtime.Time `orm:"modify_time"      json:"modify_time"`  // 更新时间
}