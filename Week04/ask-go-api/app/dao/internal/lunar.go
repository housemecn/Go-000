// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal

import (
	"database/sql"
	"time"

	"git.code.tencent.com/houseme/ask-go-api/app/model"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// LunarDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type LunarDao struct {
	gmvc.M
	Table   string
	Columns lunarColumns
}

// LunarColumns defines and stores column names for table lunar.
type lunarColumns struct {
	Id         string // ID
	Yeartime   string // 年份
	Monthtime  string // 月份
	Daytime    string // 日期
	BgImg      string // 背景图
	Weather    string // 天气内容
	Content    string // 万年历内容
	CreateTime string // 创建时间
	ModifyTime string // 修改时间
}

var (
	// Lunar is globally public accessible object for table lunar operations.
	Lunar = &LunarDao{
		M:     g.DB("default").Table("lunar").Safe(),
		Table: "lunar",
		Columns: lunarColumns{
			Id:         "id",
			Yeartime:   "yeartime",
			Monthtime:  "monthtime",
			Daytime:    "daytime",
			BgImg:      "bg_img",
			Weather:    "weather",
			Content:    "content",
			CreateTime: "create_time",
			ModifyTime: "modify_time",
		},
	}
)

// As sets an alias name for current table.
func (d *LunarDao) As(as string) *LunarDao {
	return &LunarDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *LunarDao) TX(tx *gdb.TX) *LunarDao {
	return &LunarDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *LunarDao) Master() *LunarDao {
	return &LunarDao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *LunarDao) Slave() *LunarDao {
	return &LunarDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *LunarDao) Args(args ...interface{}) *LunarDao {
	return &LunarDao{M: d.M.Args(args...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *LunarDao) LeftJoin(table ...string) *LunarDao {
	return &LunarDao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *LunarDao) RightJoin(table ...string) *LunarDao {
	return &LunarDao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *LunarDao) InnerJoin(table ...string) *LunarDao {
	return &LunarDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *LunarDao) Fields(fieldNamesOrMapStruct ...interface{}) *LunarDao {
	return &LunarDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *LunarDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *LunarDao {
	return &LunarDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *LunarDao) Option(option int) *LunarDao {
	return &LunarDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *LunarDao) OmitEmpty() *LunarDao {
	return &LunarDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *LunarDao) Filter() *LunarDao {
	return &LunarDao{M: d.M.Filter()}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (d *LunarDao) Where(where interface{}, args ...interface{}) *LunarDao {
	return &LunarDao{M: d.M.Where(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *LunarDao) WherePri(where interface{}, args ...interface{}) *LunarDao {
	return &LunarDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *LunarDao) And(where interface{}, args ...interface{}) *LunarDao {
	return &LunarDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *LunarDao) Or(where interface{}, args ...interface{}) *LunarDao {
	return &LunarDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *LunarDao) Group(groupBy string) *LunarDao {
	return &LunarDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *LunarDao) Order(orderBy ...string) *LunarDao {
	return &LunarDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *LunarDao) Limit(limit ...int) *LunarDao {
	return &LunarDao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *LunarDao) Offset(offset int) *LunarDao {
	return &LunarDao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *LunarDao) Page(page, limit int) *LunarDao {
	return &LunarDao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *LunarDao) Batch(batch int) *LunarDao {
	return &LunarDao{M: d.M.Batch(batch)}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (d *LunarDao) Cache(duration time.Duration, name ...string) *LunarDao {
	return &LunarDao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *LunarDao) Data(data ...interface{}) *LunarDao {
	return &LunarDao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.Lunar.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *LunarDao) All(where ...interface{}) ([]*model.Lunar, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.Lunar
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.Lunar.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *LunarDao) One(where ...interface{}) (*model.Lunar, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.Lunar
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *LunarDao) FindOne(where ...interface{}) (*model.Lunar, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.Lunar
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *LunarDao) FindAll(where ...interface{}) ([]*model.Lunar, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.Lunar
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Chunk iterates the table with given size and callback function.
func (d *LunarDao) Chunk(limit int, callback func(entities []*model.Lunar, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.Lunar
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *LunarDao) LockUpdate() *LunarDao {
	return &LunarDao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *LunarDao) LockShared() *LunarDao {
	return &LunarDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *LunarDao) Unscoped() *LunarDao {
	return &LunarDao{M: d.M.Unscoped()}
}