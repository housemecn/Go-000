// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"git.code.tencent.com/houseme/ask-go-api/app/dao/internal"
)

// themeMapDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type themeMapDao struct {
	*internal.ThemeMapDao
}

var (
	// ThemeMap is globally public accessible object for table {TplTableName} operations.
	ThemeMap = &themeMapDao{
		internal.ThemeMap,
	}
)

// Fill with you ideas below.