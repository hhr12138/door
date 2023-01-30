package entity

import "database/sql"

type BasicItem struct {
	Id int64 `db:"id"`
	Del sql.NullBool `db:"del"`
	GmtCreate sql.NullInt64 `db:"gmt_create"`
	GmtModified sql.NullInt64 `db:"gmt_modified"`
}
