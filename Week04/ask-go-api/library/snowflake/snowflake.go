package snowflake

import (
	"github.com/gogf/gf/frame/g"
	"github.com/housemecn/snowflake"
)

// @Project: go-pay-api
// @Author: houseme
// @Description:
// @File: snowflake
// @Version: 1.0.0
// @Date: 2020/9/22 15:54
// @Package snowflake

// InitOrderID init order id
func InitOrderID(datacenterID, workerID int64) int64 {
	s, err := snowflake.NewSnowflake(datacenterID, workerID)
	if err != nil {
		g.Log().Error(err)
		return 0
	}
	return int64(s.NextVal())
}
