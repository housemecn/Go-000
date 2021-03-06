// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id          uint64      `orm:"id,primary"          json:"id"`           // ID
	UserNo      uint64      `orm:"user_no,unique"      json:"user_no"`      // 用户ID
	Nickname    string      `orm:"nickname"            json:"nickname"`     // 昵称
	UserName    string      `orm:"user_name,unique"    json:"user_name"`    // 用户名，用户唯一标识
	UserAvatar  string      `orm:"user_avatar"         json:"user_avatar"`  // 头像
	UserSex     uint        `orm:"user_sex"            json:"user_sex"`     // 用户性别 10 男 20女 30 未知保密
	EncryptCode string      `orm:"encrypt_code,unique" json:"encrypt_code"` // 用户加密标识
	RegFrom     string      `orm:"reg_from"            json:"reg_from"`     // 邀请人的唯一标识 invite_code
	RegIp       string      `orm:"reg_ip"              json:"reg_ip"`       // 注册IP
	RegTime     *gtime.Time `orm:"reg_time"            json:"reg_time"`     // 注册时间
	RegType     uint        `orm:"reg_type"            json:"reg_type"`     // 用户注册类型 10小程序 20微信 30app
	RegOs       uint        `orm:"reg_os"              json:"reg_os"`       // 用户设备系统 10iOS、20Android、30Windows、40macOS、50其他
	UserLevel   uint        `orm:"user_level"          json:"user_level"`   // 用户登记
	AuthSalt    string      `orm:"auth_salt"           json:"auth_salt"`    // 设置密码加盐值
	Mobile      string      `orm:"mobile"              json:"mobile"`       // 用户绑定手机号
	MbTime      uint64      `orm:"mb_time"             json:"mb_time"`      // 用户绑定手机号时间
	MbState     uint        `orm:"mb_state"            json:"mb_state"`     // 手机号绑定 0 默认 100已绑定
	State       uint        `orm:"state"               json:"state"`        // 状态 100默认正常 90代审核 120禁用
	CreateTime  *gtime.Time `orm:"create_time"         json:"create_time"`  // 创建时间
	ModifyTime  *gtime.Time `orm:"modify_time"         json:"modify_time"`  // 修改时间
}
