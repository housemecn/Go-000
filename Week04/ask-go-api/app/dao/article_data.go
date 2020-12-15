// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"git.code.tencent.com/houseme/ask-go-api/app/dao/internal"
)

// articleDataDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type articleDataDao struct {
	*internal.ArticleDataDao
}

var (
	// ArticleData is globally public accessible object for table {TplTableName} operations.
	ArticleData = &articleDataDao{
		internal.ArticleData,
	}
)

// Fill with you ideas below.
